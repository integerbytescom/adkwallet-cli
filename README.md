# adkwallet-cli

adkwallet-cli is the command line wallet for ADK ( mainnet v2.1 2022 release )

adkwallet-cli can be used in two modes:

## Mode 1 - Interactive mode

You can use the adkwallet-cli executable as an interactive wallet via a character menu based interactive screen. In order to start adkwallet-cli in interactive mode, simply start the application without any parameter

When started in interactive mode, the first main screen fill look as follows:

```
***************************************************************
*
*  ADK WALLET [Network 2.1.1]  (Build 2022.03.07.1314)
*
***************************************************************

=== INIT MENU =============================
HELP   :  Show Help
META   :  METAMASK ADK Wallet [requires Metamask to be installed in Crome/Firefox/Edge]
LOCAL  :  LOCAL ADK Wallet  [stored in wallet.json]
EXIT   :  Close Wallet / Exit Program
===============================================
Enter command>
```

Here you can choose to either create/open a local wallet (stored in wallet.json), or to connect to a local Metamask installation (currently supported are Metamask extensions for Chrome, Edge and Firefox)

E.g. to create a new LOCAL wallet, enter "LOCAL" and follow the prompts:

```
=== INITWALLET MENU =============================
IMPORT   :  IMPORT account from Metamask
CREATE   :  CREATE a NEW EMPTY wallet (Generate a new SEED)
RECOVER  :  RECOVER an account by entering your 12/24 Word ADK SEED (Mnemonic)
===============================================
Enter command>
```

Once your wallet is set up or imported, <b>AND YOU HAVE SAVED THE PROVIDED RECOVERY SEED (12 word mnemonic seed) IN CASE YOU NEED TO RECOVER YOUR WALLET LATER</B>, you will find yourself on the following main wallet screen:

```
|---------- WALLET Addresses -------------------------------------------
| ( 0 ) 0x982E8ee6F3658b40d2B2DB72481d7650920A968e : 123.34 ADK
| ( 1 ) 0x53e5D435356674734DFeC856aa5a9480d2538464 : 2.123123 ADK
| ( 2 ) 0xd646F7A8b5CE24AfD8bf31720Af4B074965B17c9 : 0.0 ADK
| ( 3 ) 0x18fd4e5d73adf4Da8415935A4c010E959ef4d67C : 0.0001232 ADK
| ( 4 ) 0x73bCfC3f849D7050FEC3C286C7719140fee98542 : 0.0 ADK
|-----------------------------------------------------------------------


=== MAIN MENU =============================
BALANCE   :  Show/Refresh balances of all addresses in this wallet
SENDPOW   :  Send ADK with PoW (no Gas Fee, Proof of Work performed locally)
SENDGAS   :  Send ADK with GAS (Pay standard Gas Fees, no PoW performed)
ADD       :  Add an additional new account (derived from SEED)
QTX       :  Query a specific transaction with the TX Id
SHOWSEED  :  Shows the SEED/MNEMONIC of the current set of accounts
MIGRATE   :  Migrate OLD ADK you still hold on an v1 AZ9 SEED to this wallet
EXIT      :  Close Wallet / Exit Program
===============================================
Enter command>
```

Here you can view/refresh your wallet balance, send ADK with GAS (about 0.022 ADK to send a transaction) for INSTANT transactions, or send ADK via PoW for 0 FEE (in which case you will need to perform some PoW on your local machine (about 15s to 2 min on average).

# MIGRATE from OLD MESH:

The <b>MIGRATE</b> command lets you migrate any ADK you still hold on OLD AZ9 style addresses (in the old Mesh, on one of the old 90 character long AZ9 addresses) to your new 0x style wallet address. {Note: To do this, you will need the old SEED from your old ADK wallet (v1) - that is the old 81 character long seed also in AZ9 format}.

To start the migration process, simply select the MIGRATE command and follow the prompts. You will be asked to enter your old seed, and the new 0x destination address.



## Mode 2 - CLI Command Mode

Alternatively to the interactive mode, you can also run the adkwallet-cli program in single-command mode (non-interactive). This can be used e.g. for shell scripts or other programs that need to work with ADK, or if you just want to execute quick transactions via single shell commands.

FYI: The adkwallet-gui GUI java frontend (see https://github.com/AidosKuneen/adkwallet-gui ) uses the adkwallet-cli in command mode as the back end. All core functions are controlled by the cli (just to show you how powerful it is...)

Available commands are:

<b>ping</b> - A simple API test command, to see if you can call the API correctly
Example command:
```
adkwallet-cli.exe command ping
```
Sample response:
```
{
  "ok": true,
  "msg": "pong",
  "data": null
}
```

<b>createWalletNew</b> - creates/replaces local wallet with a brand new seed, and returns the mnemonic (to save in a secure place). The password provided will be used to encrypt the local wallet, and will be needed to open it for send transactions.
Syntax:
```
adkwallet-cli.exe command createWalletNew [new_password]
```

<b>createWalletFromMnemonic</b>  - creates/replaces local wallet from an existing mnemonic seed phrase. The password provided will be used to encrypt the local wallet, and will be needed to open it for send transactions. 
Syntax:
```
adkwallet-cli.exe command createWalletFromMnemonic ["12/24 word mnemonic seed in quotes"] [new_password]
```

<b>balance</b> - Retrieves a balance for a given 0x address (Note: the address does NOT have to be in the wallet, it can be any address)
Syntax:
```
adkwallet-cli.exe command balance [0xaddress] 
```

<b>send</b> - sends funds with GAS or with POW 
Syntax:
```
adkwallet-cli.exe command send [pow|gas] [password or "mnemonics"] [0xsource_address] [0xdestination_address] [ADK_amount]
```

<b>updatebalance</b> - updates balances for all addresses stored in wallet.json (does not need unlocking)
Syntax:
```
adkwallet-cli.exe command updatebalance 
```

<b>listwalletaddr</b> - retrieves the full list of addresses for a wallet
Syntax:
```
adkwallet-cli.exe command listwalletaddr ["mnemonic seed" or "local password"] [number_of_addresses]
```

<b>addaddress</b> - adds a new address to an existing wallet (initial default is 5 addresses, use this to add more). Note: a local wallet has to exist already
Syntax:
```
adkwallet-cli.exe command addaddress [walletpassword]
```

<b>checkpassword</b> - confirms if the password provided can be used to unlock the wallet (usually used by GUIs)
Syntax:
```
adkwallet-cli.exe command checkpassword [walletpassword]
```

<b>txinfo</b> - retrieves information about a transaction id from the node
Syntax:
```
adkwallet-cli.exe command txinfo [0xtransactionID]
```

<b>loadMetamaskMnemonics</b> - retrieves Metamask mnemonics from (Chrome/Edge/Firefox)
Syntax:
```
adkwallet-cli.exe command loadMetamaskMnemonics [password]
```

<b>migrate</b> - migrates old ADK on AZ9 addresses (from v1 SEED) to a new 0x address
Syntax: 
```
adkwallet-cli.exe command migrate [seed81char] [0xdestination]
```

Further details, examples, etc yet to be added......


