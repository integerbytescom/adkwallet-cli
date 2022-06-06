package modules

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	//"errors"
	"crypto/sha1"
	"os"
)

type Wallet struct {
	WalletName      string
	MnemonicEncoded string
	Accounts        []Account
	PasswordHash    string
	Initialized     bool
}

type Account struct {
	Id             int
	PubKey         string
	PrivKeyEncoded string
	LastBalance    string
}

var MainWallet Wallet
var MainWalletFile string
var ServerAPI string
var APIVersion string

func Sha256(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func Sha256Bytes(data []byte) string {
	h := sha1.New()
	h.Write(data)
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func Sha256Bytes2Bytes(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func InitializeConfig() {
	MainWallet = Wallet{
		WalletName:      "default",
		MnemonicEncoded: "",
		Accounts: []Account{
			Account{
				Id:             -1,
				PrivKeyEncoded: "",
				PubKey:         "",
				LastBalance:    "",
			},
		},
		PasswordHash: "",
		Initialized:  false,
	}
}

func ReadConfig() {
	if _, err := os.Stat(MainWalletFile); err == nil {
		//log.Println("reading.. ", MainWalletFile)
		jsonFile, errF := os.Open(MainWalletFile)
		if errF != nil {
			GUILog("Cannot read configuration file ", MainWalletFile, ". File corrupt or locked?")
			InitializeConfig()
		}
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			GUILog("Cannot read configuration file ", MainWalletFile, ". File corrupt or locked?")
			InitializeConfig()
		}
		err = json.Unmarshal(byteValue, &MainWallet)
		if err != nil {
			GUILog("Cannot read configuration file ", MainWalletFile, ". Content corrupted?")
			InitializeConfig()
		}
	} else { // does not exist, so create a dummy
		InitializeConfig()
	}

}

func StoreConfig() error {
	file, _ := json.MarshalIndent(MainWallet, "", " ")
	err := ioutil.WriteFile(MainWalletFile, file, 0644)
	return err
}
