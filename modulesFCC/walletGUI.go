package modules

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/term"
	"math/big"
	//"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Xmsg struct {
	Cmd string `json:"cmd"`
	Msg string `json:"msg"`
}

var Silent bool
var ForceGUILOGToStdErr bool

func init() {
	Silent = false
	ForceGUILOGToStdErr = false
}

func fmt_Println(a ...interface{}) {
	if !Silent {
		fmt.Println(a...)
	} else {
		if ForceGUILOGToStdErr {
			fmt.Fprintln(os.Stderr, a...)
		}
	}
}
func fmt_Print(a ...interface{}) {
	if !Silent {
		fmt.Print(a...)
	} else {
		if ForceGUILOGToStdErr {
			fmt.Fprint(os.Stderr, a...)
		}
	}
}

func GUISplash() {

	{
		build := "(Build 2022.04.07.4324)"
		fmt_Println("  ")
		fmt_Println("***************************************************************")
		fmt_Println("*  ")
		fmt_Println("*  ADK WALLET [Network 2.1.3] ", build)
		fmt_Println("*  ")
		fmt_Println("***************************************************************")

		time.Sleep(2 * time.Second)
	}
}

func GUIHelp() {
	{

		fmt_Println("*******************************************************************************")
		fmt_Println("Info: You can use this wallet to either connect to existing Metamask")
		fmt_Println("      accounts (Chrome/Firefox/Edge), or to open a LOCAL wallet (wallet.json)")
		fmt_Println(" ")
		fmt_Println("      If you choose 'META' in the main menu, this wallet will create a ")
		fmt_Println("      temporary connection to your Metamask account (your LOCAL wallet, if you")
		fmt_Println("      have one, will NOT be modified). It is OK to use META parallel to a local")
		fmt_Println("      wallet.")
		fmt_Println("      ")
		fmt_Println("      If you choose 'LOCAL' in the main menu, this wallet will open your ")
		fmt_Println("      local ADK wallet (stored in wallet.json) - or create one if there is none")
		fmt_Println("      ")
		fmt_Println("      To create a new local wallet, select 'LOCAL', and follow the prompts to")
		fmt_Println("      either import an account from Metamask, create a brand new wallet, or")
		fmt_Println("      recover an ADK wallet from existing mnemonics (12/24 words) ")
		fmt_Println("*******************************************************************************")
		fmt_Println("")
		fmt_Println("")

		time.Sleep(2 * time.Second)
	}
}

func GUIMenuOptions(screen string, commands []string) string {
	{

		fmt_Println("")
		fmt_Println("===", screen, "MENU =============================")
		validCommands := ""
		maxlen := Smaxlen(commands)
		for idx, cmd := range commands {
			fmt_Println(Srpad(cmd, maxlen), " : ", Lng[cmd])
			if idx != 0 {
				validCommands += " "
			}
			validCommands += cmd
		}
		fmt_Println("===============================================")
		fmt_Print("Enter command> ")
		c := "NIL"
		for {
			c = strings.ToUpper(strings.TrimSpace(ReadUserInputE()))
			fmt_Println("")
			if Scontains(commands, c) {
				return c
			}
			GUIShowError("invalid command", nil)
			GUIShowNotice("Valid options: ["+validCommands+"]", nil)
			fmt_Print("Enter command> ")
		}
	}

	return "NIL"
}

func GUIPassword() string {
	{
		fmt_Print("Enter password> ")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		input := string(bytePassword)
		fmt_Println("")
		return input
	}
	return ""
}

func GUINewPassword() string {
	{
		for {
			fmt_Print("Enter password> ")
			bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
			input := string(bytePassword)
			if len(input) < 8 {
				fmt_Println("Error: password too short")
				continue
			}
			fmt_Print("\nEnter password again> ")
			bytePassword2, _ := term.ReadPassword(int(syscall.Stdin))
			input2 := string(bytePassword2)
			if input2 != input {
				fmt_Println("Error: passwords dont match")
				continue
			}
			return input
		}
	}
	return ""
}

func GUIShowMnemonic(newMM string) {
	{
		fmt_Println("")
		fmt_Println("=== SEED / MNEMONIC ============================================================")
		fmt_Println("  " + newMM)
		fmt_Println("================================================================================")
		fmt_Println(" Note: Store this SEED in a safe place. ")
		fmt_Println(" This seed will allow you to recover your wallet in case you forget your  ")
		fmt_Println(" password or if your wallet file gets corrupted. Do not share the seed ")
		fmt_Println(" with anyone! ")
		fmt_Println("================================================================================")
		fmt_Println("")
	}
}

func GUIShowAccount() {
	{
		fmt_Println("")
		fmt_Println("|---------- WALLET Addresses -------------------------------------------")
		for id, account := range MainWallet.Accounts {
			fmt_Println("|", "(", id, ")", account.PubKey, ":", printWEIasADK(account.LastBalance, true))
		}
		fmt_Println("|-----------------------------------------------------------------------")
		fmt_Println("")
	}

}

func GUIChooseFromMM(mms []string) string {
	accountlist := []string{}
	for _, mm := range mms {
		firstAccount := GetAccountForMM(mm, 0)
		accountlist = append(accountlist, firstAccount)
	}

	{
		fmt_Println("")
		fmt_Println("==== Connected Metamask Accounts =====================================")
		for idx, acct := range accountlist {
			fmt_Println("(", idx+1, ") ", acct, " [...] ")
		}
		fmt_Println("======================================================================")
		for {
			fmt_Println("Please enter the number of the account you want to open [", 1, "-", len(accountlist), "]:")
			fmt_Print("account number> ")
			intVar, err := strconv.Atoi(ReadUserInputE())
			if err == nil && intVar >= 1 && intVar <= (len(accountlist)) {
				return mms[intVar-1]
			}
		}
	}
	return ""
}

func GUIGetTX() (string, bool) {
	{
		tx := GUIGetTextInput("TX", "enter TX hash")
		//
		match, _ := regexp.MatchString("^0x([A-Fa-f0-9]{64})$", tx)
		if match {
			return tx, true
		}
		return "", false
	}

	return "", false
}

func GUIGetAZ9SEED() (string, bool) {
	{
		tx := GUIGetTextInput("AZ9SEED", "Paste your OLD 81 character long AZ9 seed (old v1 wallet seed)\n")
		//
		match, _ := regexp.MatchString("^([A-Z9]{81})$", tx)
		if match {
			return tx, true
		}
		return "", false
	}

	return "", false
}

func GUIGetDestinationAddress() (string, bool) {

	{
		addr := GUIGetTextInput("ADDRESS", "enter 0x destination address")
		//
		valid := IsValidAddress(addr)
		if valid {
			return addr, true
		}
		return "", false
	}

	return "", false
}

func GUIReadyToSend(sendfrom string, amount_bigf big.Float, destination string) bool {
	{
		s_msg := ("***********************************\n") +
			("** READY TO SEND ?\n") +
			("** From: " + sendfrom + "\n") +
			("** To: " + destination + "\n") +
			("** Amount: " + removeTrailing0(amount_bigf.Text('f', 18)) + " ADK\n") +
			("***********************************\n")

		return GUIConfirmYesNo(s_msg)
	}
	return false
}

func GUIStartPOW() {
	{
		fmt_Println("starting PoW (Proof of Work). This can take a while...")
	}
}

func GUIStopPOW() {
	{
		fmt_Println("PoW complete.")
	}
}

func GUIWaitingMining() {
	{
		fmt_Println("Waiting for Transaction to be mined...")
	}
}

func GUIConfirmYesNo(message string) bool {
	{
		for {
			var input string
			fmt_Print(message + "(Enter Yes or No ): ")
			input = ReadUserInputE()
			input = strings.ToLower(input)
			if input == "n" || input == "no" {
				return false
			} else if input == "y" || input == "yes" {
				return true
			}
		}
	}
	return false
}

func GUIShowTXInfo(tx string) error {
	{
		txdata, from, receiptjson, err := (GetTXInfo(tx))
		if err != nil {
			return err
		}
		if txdata == nil {
			return errors.New("empty transaction data")
		}
		fmt_Println("")
		fmt_Println("================TRANSACTION INFO =============== ")
		fmt_Println(" Hash: ", txdata.Hash)
		fmt_Println(" Value (WEI): ", txdata.Value)
		fmt_Println(" Value (ADK): ", printWEIasADK(txdata.Value, true))
		fmt_Println(" Gas: ", txdata.Gas)
		fmt_Println(" GasPrice: ", txdata.GasPrice)
		fmt_Println(" From: ", from)
		fmt_Println(" To: ", txdata.To)
		fmt_Println(" Pending: ", txdata.Pending)
		fmt_Println(" Nonce: ", txdata.Nonce)
		fmt_Println(" RECEIPT: ", receiptjson)
		fmt_Println("================================================== ")
		fmt_Println("")

		return nil
	}

	return nil
}

func GUIGetExistingAccount() (int, error) {
	{
		lastAcctId := strconv.Itoa(len(MainWallet.Accounts) - 1)
		acctno_s := GUIGetTextInput("SENDFROM", "Enter address (number) to SEND FROM ( 0 - "+lastAcctId+" or STOP to abort )")
		if strings.ToUpper(acctno_s) == "STOP" {
			return -1, errors.New("cancelled by user")
		}
		acctno, err := strconv.Atoi(acctno_s)
		if err != nil || acctno < 0 || acctno > len(MainWallet.Accounts)-1 {
			return -1, errors.New("Invalid wallet selected")
		}
		return acctno, nil
	}
	return -1, nil
}

func GUIGetSendAmount(maxAmount string, sendPOW bool) (big.Float, error) {

	f0, _, _ := big.ParseFloat("0.021", 10, 256, big.ToNearestEven)

	{
		amount_in_ADK := GUIGetTextInput("ADK", "Enter ADK amount so send, or STOP to abort (max. "+printWEIasADK(maxAmount, true)+")")
		if strings.ToUpper(amount_in_ADK) == "STOP" {
			return *f0, errors.New("user aborted")
		}

		if !IsValidAmount(amount_in_ADK) {
			return *f0, errors.New("Invalid amount entered, check your number format. Decimal separator is DOT (e.g.  12.3 )")
		}

		if !strings.Contains(amount_in_ADK, ".") {
			amount_in_ADK += ".0"
		}

		amount_bigf, _, _ := big.ParseFloat(amount_in_ADK, 10, 256, big.ToNearestEven)
		amount_bigf_avail, _, _ := big.ParseFloat(printWEIasADK(maxAmount, false), 10, 256, big.ToNearestEven)

		if !sendPOW {
			fees, _, _ := big.ParseFloat("0.021", 10, 256, big.ToNearestEven)
			amount_bigf_avail.Sub(amount_bigf_avail, fees)
			// deduct 0.021 for gas fees
		}

		enough := amount_bigf.Cmp(amount_bigf_avail) <= 0

		if !enough && sendPOW {
			return *f0, errors.New("insufficient balance available on selected address")
		}

		if !enough && !sendPOW {
			return *f0, errors.New("insufficient address balance (amount entered plus GAS fee of 0.021 ADK exceeds balance)")
		}

		return *amount_bigf, nil

	}

	return *f0, nil
}

func IsValidAmount(v string) bool {
	re := regexp.MustCompile("^[0-9]{1,8}([.][0-9]{1,18})?$")
	return re.MatchString(v)
}

func GUIGetTextInput(inputtype string, label string) string {
	_ = inputtype // TBD
	{
		fmt_Print(label + "> ")
		reader := bufio.NewReader(os.Stdin)
		input := ""
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)
		return strings.TrimSpace(input)
	}
	return ""
}

func GUIShowNotice(msg string, err error) {
	{
		if err != nil {
			fmt_Println("** Notice:", msg, " ", err)
		} else {
			fmt_Println("** Notice:", msg)
		}
	}
}

func GUILog(a ...interface{}) {
	{
		fmt_Println("** Log: ", a)
	}
}

func GUIPrint(s string) {
	{
		fmt_Print(".")
	}
}

func GUIPrintErr(s string) {
	// this we do print  if silent.. its for the POW status to the GUI Wallet
	if Silent {
		fmt.Fprintf(os.Stderr, s)
	}
}

func GUIShowError(msg string, err error) {
	{
		if err != nil {
			fmt_Println("!!!!!")
			fmt_Println("!! Error:", msg, " ", err)
			fmt_Println("!!!!!")
		} else {
			fmt_Println("!!!!!")
			fmt_Println("!! Error:", msg)
			fmt_Println("!!!!!")
		}
	}
}

func GUIShowFatal(msg string, err error) {
	{
		fmt_Println("** FATAL:", msg, err)
	}
	os.Exit(1)
}

// helper
func Scontains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func Smaxlen(s []string) int {
	l := 0
	for _, v := range s {
		if len(v) > l {
			l = len(v)
		}
	}
	return l
}

func Srpad(s string, lg int) string {
	if len(s) > lg {
		return s[:lg]
	}
	ret := s
	for {
		if len(ret) == lg {
			return ret
		}
		ret += " "
	}
}

func IsValidAddress(v string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(v)
}

func removeTrailing0(str string) string {
	if !strings.Contains(str, ".") {
		return str
	}
	v_ret := str
	for {
		if v_ret[len(v_ret)-1:] == "0" && v_ret[len(v_ret)-2:len(v_ret)-1] != "." {
			v_ret = v_ret[:len(v_ret)-1]
		} else {
			break
		}
	}
	return v_ret
}

func printWEIasADK(val string, curr bool) string {
	val_new := strings.TrimSpace("00000000000000000000" + val)
	val_right := val_new[len(val_new)-18:]
	val_left := val_new[:len(val_new)-18]
	val_new = val_left + "." + val_right
	for { //remove leading 0
		if val_new[0:1] == "0" && val_new[1:2] != "." {
			val_new = val_new[1:]
		} else {
			break
		}
	}
	for { //remove trailing 0
		if val_new[len(val_new)-1:] == "0" && val_new[len(val_new)-2:len(val_new)-1] != "." {
			val_new = val_new[:len(val_new)-1]
		} else {
			break
		}
	}
	if curr {
		return val_new + " ADK"
	} else {
		return val_new
	}
}
