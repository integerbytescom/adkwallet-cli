package modules

import (
	"errors"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"math/big"
)

var IsGui bool
var Lng map[string]string

func init() {
	IsGui = false
	Lng = make(map[string]string)

	Lng["HELP"] = "Show Help"
	Lng["META"] = "METAMASK ADK Wallet [requires Metamask to be installed in Crome/Firefox/Edge]"
	Lng["LOCAL"] = "LOCAL ADK Wallet  [stored in wallet.json]"

	Lng["BALANCE"] = "Show/Refresh balances of all addresses in this wallet"
	Lng["SENDPOW"] = "Send ADK with PoW (no Gas Fee, Proof of Work performed locally)"
	Lng["SENDGAS"] = "Send ADK with GAS (Pay standard Gas Fees, no PoW performed)"
	Lng["ADD"] = "Add an additional new account (derived from SEED)"
	Lng["QTX"] = "Query a specific transaction with the TX Id"
	Lng["MIGRATE"] = "Migrate OLD ADK you still hold on an v1 AZ9 SEED to this wallet"
	Lng["SHOWSEED"] = "Shows the SEED/MNEMONIC of the current set of accounts"
	Lng["EXIT"] = "Close Wallet / Exit Program"

	Lng["IMPORT"] = "IMPORT account from Metamask"
	Lng["CREATE"] = "CREATE a NEW EMPTY wallet (Generate a new SEED)"
	Lng["RECOVER"] = "RECOVER an account by entering your 12/24 Word ADK SEED (Mnemonic)"
}

func OpenWallet_Flow() error {
	GUISplash()
	return OpenMainMenu_Flow()
}
func OpenMainMenu_Flow() error {

	switch GUIMenuOptions("INIT", []string{"HELP", "META", "LOCAL", "EXIT"}) {
	case "HELP":
		GUIHelp()
		return OpenMainMenu_Flow()

	case "EXIT":
			StoreConfig()
			return nil

	case "META":
		//
		GUIShowNotice("Connecting to METAMASK... Please enter your METAMASK password", nil)

		pw := GUIPassword()
		GUIShowNotice("Connecting to METAMASK (Chrome/Firefox/Edge) with the provided credentials...", nil)
		mms := ScanForMnemonics(pw) // check if this is Metamask or Local, or both

		if len(mms) <= 0 {
			GUIShowError("INVALID METAMASK PASSWORD or METAMASK NOT INSTALLED.", nil)
			OpenMainMenu_Flow()
		} else {
			mm := GUIChooseFromMM(mms)
			MainWalletFile = "tmplink.json" // only temporary
			InitializeLocalWalletFromMM(mm, pw)
		}
		return MainWallet_Flow()
	case "LOCAL":
		if MainWallet.Initialized { // local wallet exists
			GUIShowNotice("opening local wallet...", nil)
			for {
				pw := GUIPassword()
				if IsValidLocalPassword(pw) {
					CachedPassword = pw
					return MainWallet_Flow()
				} else {
					GUIShowError("password invalid", nil)
				}
			}
		} else { // NOT MainWallet.Initialized
			InitializeWallet_Flow()
			return MainWallet_Flow()
		}
	} //end switch
	return nil
}

func MainWallet_Flow() error {

	GUIShowNotice("opening wallet...", nil)
	GUIShowNotice("updating balances from node. please wait", nil)
	errBal1 := UpdateBalances()
	if errBal1 != nil {
		GUIShowError("Unable to update balances", errBal1)
	}
	GUIShowAccount()

	for {
		StoreConfig()
		switch GUIMenuOptions("MAIN", []string{"BALANCE", "SENDPOW", "SENDGAS", "ADD", "QTX", "SHOWSEED", "MIGRATE", "EXIT"}) {
		case "BALANCE":
			GUIShowNotice("updating balances from node. please wait", nil)
			errBal := UpdateBalances()
			if errBal != nil {
				GUIShowError("Unable to update balances", errBal)
			} else {
				StoreConfig()
			}
			GUIShowAccount()
		case "SENDPOW":
			MainWalletSend_Flow(true)
			StoreConfig()
		case "SENDGAS":
			MainWalletSend_Flow(false)
			StoreConfig()
		case "ADD":
			addr, errAddr := AddAddress()
			if errAddr != nil {
				GUIShowError("Unable to add address", errAddr)
			} else {
				GUIShowNotice("address added: "+addr, nil)
				GUIShowNotice("updating balances from node. please wait", nil)
				UpdateBalances()
				GUIShowAccount()
				StoreConfig()
			}
		case "QTX":
			tx, valid := GUIGetTX()
			if valid {
				err := GUIShowTXInfo(tx)
				if err != nil {
					GUIShowError("Unable to load transaction info", err)
				}
			} else {
				GUIShowError("invalid transaction hash.", nil)
			}
		case "SHOWSEED":
			mm, _ := DecryptString(CachedPassword, MainWallet.MnemonicEncoded)
			GUIShowMnemonic(mm)
		case "MIGRATE":
			// migrating old AZ9
			GUIShowNotice("Select the DESTINATION 0x ADDRESS to send migrated ADK to.", nil)
			destAddress, validDest := GUIGetDestinationAddress()
			if !validDest {
				GUIShowError("Destination address error. Not a valid 0x address.", nil)
			} else {
				az9seed, valid := GUIGetAZ9SEED()
				if !valid || !IsValidAddress(destAddress) {
					GUIShowError("invalid seed. must be 81 charcters long and only consist of the letters A-Z and 9.", nil)
				} else {
					GUIMigrate_Flow(destAddress, az9seed, 10000)
				}

			}

		case "EXIT":
			StoreConfig()
			return nil
		}
	}
	StoreConfig()
	return nil
}

func MainWalletSend_Flow(usePOW bool) {

	accountID, aerr := GUIGetExistingAccount()
	if aerr != nil {
		GUIShowError("Error selecting wallet", aerr)
	} else {
		GUIShowNotice("updating balances from node. please wait", nil)
		UpdateBalances()
		amount_bigf, errA := GUIGetSendAmount(MainWallet.Accounts[accountID].LastBalance, usePOW)
		if errA != nil {
			GUIShowError("Amount error", errA)
		} else {
			destination, valid := GUIGetDestinationAddress()
			if !valid {
				GUIShowError("Invalid destination address entered", nil)
			} else {
				vOK := GUIReadyToSend(MainWallet.Accounts[accountID].PubKey, amount_bigf, destination)
				if !vOK {
					GUIShowError("aborted by user", nil)
				} else {
					// ready to send
					privateKey_str, _ := DecryptString(CachedPassword, MainWallet.Accounts[accountID].PrivKeyEncoded)
					amount_bigf_WEI, _, _ := big.ParseFloat("0", 10, 256, big.ToNearestEven)
					amount_bigf_18, _, _ := big.ParseFloat("1000000000000000000", 10, 256, big.ToNearestEven)
					amount_bigf_WEI.Mul(&amount_bigf, amount_bigf_18)
					amount_bigInt_WEI := new(big.Int)
					amount_bigf_WEI.Int(amount_bigInt_WEI)

					var tx string
					var errtx error
					if usePOW {
						tx, errtx = TransferADKWithPoW(privateKey_str, destination, amount_bigInt_WEI, []byte{0,0,0,0})
					} else {
						tx, errtx = TransferADKWithGas(privateKey_str, destination, amount_bigInt_WEI, 21000)
					}
					if errtx == nil {
						GUIShowNotice("TX SENT AND MINED: "+tx, nil)
						GUIShowNotice("updating balances in local wallet...", nil)
						UpdateBalances()
					} else {
						GUIShowError("Error while sending. ", errtx)
					}
				}
			}
		}
	}
}

func IsValidLocalPassword(p string) bool {
	return (Sha256(p+"salt628374") == MainWallet.PasswordHash)
}

func InitializeWallet_Flow() error {

	GUIShowNotice("NEW WALLET SETUP: Setting up LOCAL wallet (as none exists currently)", nil)

	switch GUIMenuOptions("INITWALLET", []string{"IMPORT", "CREATE", "RECOVER"}) {
	case "IMPORT":
		GUIShowNotice("Connecting to Metamask...", nil)
		password := GUIPassword()
		mms := ScanForMnemonics(password)
		if len(mms) > 0 {
			mm := GUIChooseFromMM(mms)
			InitializeLocalWalletFromMM(mm, password)
		} else {
			GUIShowError("Could not connect to a valid Metamask account.\nPlease check if Metamask is installed [Chrome/Firefox] and then try again.", nil)
			return errors.New("invalid password")
		}
	case "CREATE":
		password := ""
		GUIShowNotice("Creating a new encrypted wallet. Please choose a new password (min 8 char)", nil)
		password = GUINewPassword()
		newMM, _ := hdwallet.NewMnemonic(128)
		GUIShowNotice("NEW Wallet created. Please save your Seed/Mnemonic words in a safe place.", nil)
		GUIShowMnemonic(newMM)
		InitializeLocalWalletFromMM(newMM, password)


	case "RECOVER":
		password := ""
		GUIShowNotice("Recovering your wallet from a known SEED / Mnemonic (12 or 24 words)", nil)

		mmseed := GUIGetTextInput("MNEMONIC", "Please enter/paste your 12 or 24 words (Mnemonic) SEED\n")
		pk0 := GetAccountForMM(mmseed, 0)
		if len(pk0) < 20 {
			GUIShowError("Not a valid 12/24 word seed.", nil)
			return errors.New("Not a valid 12/24 word seed.")
		}
		GUIShowNotice("Wallet recovered", nil)
		GUIShowNotice("Choose a password to encrypt your wallet file", nil)
		password = GUINewPassword()

		InitializeLocalWalletFromMM(mmseed, password)


	}

	return nil
}
