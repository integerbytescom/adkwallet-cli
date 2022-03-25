package modules

import (
	"bytes"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"strconv"
	"strings"
)

var (
	kBuf, vBuf bytes.Buffer
)

var CachedPassword string
var Authenticated bool
var loadAccounts int

func init() {
	CachedPassword = ""
	Authenticated = true
	loadAccounts = 5
}

func InitializeLocalWalletFromMM(mm string, password string) {
	CachedPassword = password
	MainWallet = Wallet{
		WalletName:      "default",
		MnemonicEncoded: EncryptString(password, mm),
		Accounts:        []Account{},
		PasswordHash:    Sha256(password + "salt628374"),
		Initialized:     true,
	}

	for i := 0; i < loadAccounts; i++ {
		MainWallet.Accounts = append(MainWallet.Accounts,
			Account{
				Id:             i,
				PrivKeyEncoded: EncryptString(password, GetAccountPKForMM(mm, i)),
				PubKey:         GetAccountForMM(mm, i),
			})
	}

}

func AddAddress() (string, error) {
	lastID := -1
	for id, _ := range MainWallet.Accounts {
		lastID = id
	}
	lastID++
	mm, err := DecryptString(CachedPassword, MainWallet.MnemonicEncoded)
	if err != nil {
		return "", err
	}
	newaddr := GetAccountForMM(mm, lastID)
	MainWallet.Accounts = append(MainWallet.Accounts,
		Account{
			Id:             lastID,
			PrivKeyEncoded: EncryptString(CachedPassword, GetAccountPKForMM(mm, lastID)),
			PubKey:         newaddr,
		})
	loadAccounts = len(MainWallet.Accounts)
	return newaddr, nil
}

func UpdateBalances() error {
	client, err := ethclient.Dial(ServerAPI)
	if err != nil {
		return err
	}
	for id, account := range MainWallet.Accounts {
		bal, errbal1 := GetAddressBalance(*client, account.PubKey)
		if errbal1 != nil {
			return errbal1
		} else {
			if bal == "" {
				bal = "0"
			}
			MainWallet.Accounts[id].LastBalance = bal
		}
	}
	return nil
}

func AddressBalances(addrs []string) ([]string, error) {
	if len(addrs) == 0 {
		return nil, errors.New("invalid address: nil")
	}
	for _, addr := range addrs {
		if !IsValidAddress(addr) {
			return nil, errors.New("invalid address: " + addr)
		}
	}

	ret := []string{}

	client, err := ethclient.Dial(ServerAPI)
	if err != nil {
		return ret, err
	}

	for _, addr := range addrs {
		bal, errbal1 := GetAddressBalance(*client, addr)
		if errbal1 != nil {
			ret = append(ret, "0")
			return nil, errors.New("error fetching balance for: " + addr)

		} else {
			if bal == "" {
				bal = "0"
			}
			ret = append(ret, bal)
		}
	}
	return ret, nil
}

func GetAccountForMM(mm string, id int) string { //id = account number, starting with 0
	wallet, err := hdwallet.NewFromMnemonic(strings.TrimSpace(mm))
	if err != nil {
		return "NOT VALID"
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.Itoa(id))
	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")//
	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")//
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "NOT VALID"
	}
	return account.Address.Hex()
}

func GetAccountPKForMM(mm string, id int) string { //id = account number, starting with 0
	wallet, err := hdwallet.NewFromMnemonic(strings.TrimSpace(mm))
	if err != nil {
		return "NOT VALID"
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.Itoa(id))
	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")//
	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")//

	account, err := wallet.Derive(path, false)
	pk_hx, err2 := wallet.PrivateKeyHex(account)

	if err != nil || err2 != nil {
		return "NOT VALID"
	}
	return pk_hx
}
