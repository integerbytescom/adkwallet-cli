package main

import (
	"fmt"
	modules "github.com/AidosKuneen/adkWallet/modules"
	"log"
	"os"
	"strings"
)
var HaveCommand bool
func main() {

 	modules.MainWalletFile = "wallet.json" //initially we use this file, but if user logs on to metamask we use a different one
	modules.APIVersion = "2"               // api version'

	args := os.Args[1:]


	HaveCommand := false
	commandline := []string{}
	//
	var apiFilename = modules.MainWalletFile
	// check for adknode parameter
	modules.ServerAPI = "http://api1.mainnet.aidoskuneen.com:9545"
	for idx, arg := range args {
		if (strings.ToLower(arg) == "-adknode" ||
			strings.ToLower(arg) == "adknode") && idx < len(arg)-2 {
			modules.ServerAPI = args[idx+1]
		}
		if (strings.ToLower(arg) == "-apiwallet" ||
			strings.ToLower(arg) == "apiwallet") && idx < len(arg)-2 {
			apiFilename = args[idx+1]
		}
		if (strings.ToLower(arg) == "-apiversion" ||
			strings.ToLower(arg) == "apiversion") && idx < len(arg)-2 {
			if args[idx+1] != modules.APIVersion {
				log.Fatal("API VERSION INVALID. requested: ", args[idx+1], " actual:", modules.APIVersion, "Ensure you use the most up to date versions")
			}
		}
		if (strings.ToLower(arg) == "adk-wallet-daemon" && os.Getenv("ADKmonitor")=="true"){
			  modules.Silent = true
			  go modules.OpenWallet_Flow() //run as daemon to keep wallet up to date
		}
	}
  // check for command parameter
	for _, arg := range args {
		if HaveCommand { // all after command keyword
			commandline = append(commandline, arg)
		}
		if arg == "command" || arg == "-command" {
			HaveCommand = true
			modules.MainWalletFile = apiFilename // for api calls, we use the cmd line file name. Param is not accepted for text GUI use
		}
	}


	modules.ReadConfig()
	modules.StoreConfig()

	if HaveCommand {
		modules.Silent = true
		fmt.Println(modules.ExecuteCommand(commandline))
	} else {
		modules.OpenWallet_Flow()
	}

}
