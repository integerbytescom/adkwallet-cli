package modules

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	Models "github.com/AidosKuneen/adkWallet/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"log"
)

func GetLatestBlockNumber(client *ethclient.Client) (string) {
		 header, err := client.HeaderByNumber(context.Background(), nil)
		 if err != nil {
				 log.Println(err)
				 return "0"
		 }

		 return header.Number.String() // 5671744
}

// GetLatestBlock from blockchain
func GetLatestBlock(client ethclient.Client) *Models.Block {
	// We add a recover function from panics to prevent our API from crashing due to an unexpected error
	defer func() {
		if err := recover(); err != nil {
			//fmt.PXrintln(err)
		}
	}()

	// Query the latest block
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		GUILog(err)
		return nil
	}

	// Build the response to our model
	_block := &Models.Block{
		BlockNumber:       block.Number().Int64(),
		Timestamp:         block.Time(),
		Difficulty:        block.Difficulty().Uint64(),
		Hash:              block.Hash().String(),
		TransactionsCount: len(block.Transactions()),
		Transactions:      []Models.Transaction{},
	}

	for _, tx := range block.Transactions() {
		_block.Transactions = append(_block.Transactions, Models.Transaction{
			Hash:     tx.Hash().String(),
			Value:    tx.Value().String(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice().Uint64(),
			Nonce:    tx.Nonce(),
			To:       tx.To().String(),
		})
	}

	return _block
}

// GetTxByHash by a given hash
func GetTxByHash(client ethclient.Client, hash common.Hash) *Models.Transaction {

	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		//fmt.PXrintln(err)
	}

	return &Models.Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		To:       tx.To().String(),
		Pending:  pending,
		Nonce:    tx.Nonce(),
	}
}

func GetTXInfo(hashStr string) (*Models.Transaction, string, string, error) {

	client, errC := ethclient.Dial(ServerAPI)
	if errC != nil {
		return nil, "", "", errC
	}
	hash := common.HexToHash(hashStr)
	tx, pending, errH := client.TransactionByHash(context.Background(), hash)
	if errH != nil {
		return nil, "", "", errH
	}

	from := ""
	if msg, err := tx.AsMessage(types.NewEIP155Signer(big.NewInt(40272)), nil); err == nil {
		from = msg.From().Hex()
	}

	receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
	b_receipt_json, _ := json.MarshalIndent(receipt, " ", " ")

	return &Models.Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		To:       tx.To().String(),
		Pending:  pending,
		Nonce:    tx.Nonce(),
	}, from, string(b_receipt_json), nil

}

// TransferEth from one account to another
func TransferADKWithGas(privKey string, to string, amount *big.Int, i_gaslimit int) (string, error) {
	client, errC := ethclient.Dial(ServerAPI)
	if errC != nil {
		return "", errC
	}
	// Assuming you've already connected a client, the next step is to load your private key.
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	// Function requires the public address of the account we're sending from -- which we can derive from the private key.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Now we can read the nonce that we should use for the account's transaction.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	var data []byte
	value := (amount)                     // in wei
	gasLimit := uint64(i_gaslimit)             // in units, 21000 or more
	gasPrice := big.NewInt(1000000000000) //  always 1000 Gwei for ADK - client.SuggestGasPrice(context.Background())

	// We figure out who we're sending the ADK to.
	toAddress := common.HexToAddress(to)
	// We create the transaction payload
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID := big.NewInt(40272) // ADK PROD

	// We sign the transaction using the sender's private key
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// Now we are finally ready to broadcast the transaction to the entire network
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	GUIWaitingMining()

	receipt, errB := bind.WaitMined(context.Background(), client, signedTx)

	if errB != nil {
		return signedTx.Hash().String(), errB
	}
	if receipt.Status != 1 { // reverted, extract the revert reason
		return signedTx.Hash().String(), errors.New("TRANSACTION REVERTED")
	}

	// We return the transaction hash
	return signedTx.Hash().String(), nil
}

// TransferEth from one account to another
func TransferADKWithPoW(privKey string, to string, amount *big.Int, functioncall []byte) (string, error) {

	client, errC := ethclient.Dial(ServerAPI)
	if errC != nil {
		return "", errC
	}
	// Assuming you've already connected a client, the next step is to load your private key.
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	// Function requires the public address of the account we're sending from -- which we can derive from the private key.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Now we can read the nonce that we should use for the account's transaction.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := (amount) // in wei
	//gasLimit := uint64(21000)   // in units
	gasLimit := uint64(500000) // in units // since its free we can set it higher, also needs a bit higher since DATA field is not blank
	gasPrice := big.NewInt(0)  //  always 0 Gwei for ADK -  if Pow done locally

	// We figure out who we're sending the ADK to.
	toAddress := common.HexToAddress(to)

	GUIStartPOW()
	GUIPrintErr("PoW:")
	// We create the transaction payload
	var tx *types.Transaction
	var signedTx *types.Transaction
	var cnt int64
	cnt = 0
	for {
		var token []byte
		token = make([]byte, 32)
		rand.Read(token) // random, but first 4 bytes are always 0
		if len(functioncall) < 4 {
			token[0] = 0
			token[1] = 0
			token[2] = 0
			token[3] = 0
		} else {
			token[0] = functioncall[0]
			token[1] = functioncall[1]
			token[2] =  functioncall[2]
			token[3] =  functioncall[3]
		}
		tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, token)
		chainID := big.NewInt(40272) // ADK PROD
		// We sign the transaction using the sender's private key
		signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return "", err
		}
		//
		PoWHash := Sha256Bytes2Bytes(signedTx.Hash().Bytes())
		//
		if PoWHash[0] == 0 &&
			PoWHash[1] == 0 &&
			PoWHash[2] < 16 { //64
			break
		}
		cnt++
		if cnt%1000 == 0 {
			GUIPrint(".")
		} // for CLI
		if (cnt % 4000) == 0 {
			GUIPrintErr(".")
		} // this is so we still see it in the log output for the GUI
		if ((cnt + 1000) % 4000) == 0 {
			GUIPrintErr("-")
		} // this is so we still see it in the log output for the GUI
		if ((cnt + 2000) % 4000) == 0 {
			GUIPrintErr("'")
		} // this is so we still see it in the log output for the GUI
		if ((cnt + 3000) % 4000) == 0 {
			GUIPrintErr("-")
		} // this is so we still see it in the log output for the GUI
		if (cnt % 100000) == 0 {
			GUIPrintErr("\n")
		}
		// if (cnt%100000)==0 {
		// 	fmt.PXrintln("\n"+signedTx.Hash().String()[:10])
		// }
	}

	GUIStopPOW()
	GUIPrintErr("\npow done. waiting for transaction to be mined...")

	// Now we are finally ready to broadcast the transaction to the entire network
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	GUIWaitingMining()

	receipt, errB := bind.WaitMined(context.Background(), client, signedTx)

	if errB != nil {
		return signedTx.Hash().String(), errB
	}
	if receipt.Status != 1 { // reverted, extract the revert reason
		return signedTx.Hash().String(), errors.New("TRANSACTION REVERTED")
	}
	// We return the transaction hash
	return signedTx.Hash().String(), nil
}

// GetAddressBalance returns the given address balance =P
func GetAddressBalance(client ethclient.Client, address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "0", err
	}

	return balance.String(), nil
}


func GetAuthUnstake(client *ethclient.Client, s_privateKey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(s_privateKey)
	if err != nil {
		log.Println(err)
		return nil
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return nil
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return nil
	}

	chainID := big.NewInt(40272)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(210000) //
	auth.GasPrice = big.NewInt(1000000000000)//  always 1000 Gwei for ADK
	return auth
}
