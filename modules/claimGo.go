package modules

// CLAIM MODULE for old AZ9 Accounts with ADKs in them

import (
	"context"
	"crypto/ecdsa"
	"github.com/AidosKuneen/gadk"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var adkTransactionsContract common.Address

var adk *ADKTransactions
var initialized bool //default is false

func GetAuthClaim(client *ethclient.Client) *bind.TransactOpts {
	// the generic API client private key is 0x8ddda563583494672352748957abccceff773867dafe5187263541827ffaee8f
	// the generic API client public key is 0x2B5f3EC809994eD4549d4305fCf430129Dd96A3D
	// this is OK to be PUBLIC as this account is used for PoW validation, and not to keep any ADK

	// we use a KNOWN private key, and it is OK this key is known. It is not used for anything but to submit MESH transactions
	privateKey, err := crypto.HexToECDSA("8ddda563583494672352748957abccceff773867dafe5187263541827ffaee8f")

	if err != nil {
		GUILog(err)
		return nil
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		GUILog("error casting public key to ECDSA")
		return nil
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		GUILog(err)
		return nil
	}

	chainID := big.NewInt(40272)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei

	auth.GasLimit = (GetGasLimit(client) / 10) * 9 // 90% of actual gas limit from last block

	auth.GasPrice = big.NewInt(0)
	return auth
}

func GetBlockNumber(client *ethclient.Client) int64 {
	GUILog("GetBlockNumber") //
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		GUILog(err)
		return -1
	}
	return header.Number.Int64()
}

func GetGasLimit(client *ethclient.Client) uint64 {
	GUILog("GasLimit in LastBlock: ") //
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		GUILog(err)
		return 450000000
	}
	return header.GasLimit
}

func toByte32(s []byte) [32]byte {
	ret := [32]byte{}
	if len(s) >= 32 {
		for i := 0; i < 32; i++ {
			ret[i] = s[i]
		}
	}
	return ret
}

type RichList1 struct {
	Addresses      []string `json:"addresses"`
	Balances       []string `json:"balances"`
	Milestone      string   `json:"milestone"`
	MilestoneIndex int64    `json:"milestoneIndex"`
	Duration       int64    `json:"duration"`
}

// IntPow calculates n to the mth power. Since the result is an int, it is assumed that m is a positive power
func IntPow(n, m int64) int64 {
	if m == 0 {
		return 1
	}
	result := n
	var i int64
	for i = 2; i <= m; i++ {
		result *= n
	}
	return result
}

func GUIMigrate_Flow(p_targetAirdrop string, p_seed string, countAddresses int) (*big.Int, bool) {
	successfullyClaimed := big.NewInt(0)
	seed := &p_seed
	targetAirdrop := &p_targetAirdrop
	var richList = Richlist
	GUILog("Checking totals from genesis 2.1 snapshot...")
	//
	addressToValueMap := make(map[string]int64)
	for i, addr := range richList.Addresses {
		val, err := strconv.ParseInt(richList.Balances[i], 10, 64)
		if err == nil {
			addressToValueMap[addr] += val
		} else {
			GUILog("Error parsing balance: ", addr, richList.Balances[i])
			return successfullyClaimed, false
		}
	}
	// check totals -  must contain entire ADK balance
	var total int64
	total = 0

	var addresses_as_array []string

	a_idx := 0
	for _ad, _val := range addressToValueMap {
		addresses_as_array = append(addresses_as_array, _ad)
		total += _val
		a_idx++
	}

	if total != 887167045722474 { // note 887167045722474 is what's remeining after v1-2 migration in genesis 2.1
		GUILog("Total is not 8871670.45722474 ADK. Are you using the correct GENESIS milestone json file? Something's wrong: ", total)
		return successfullyClaimed, false
	}
	GUILog("Connecting to ADKgo node db " + ServerAPI)
	client, err := ethclient.Dial(ServerAPI)
	GUILog("Block: ", GetBlockNumber(client))

	GUILog("Scanning seed for available balances to claim... (this can take a while)")

	foundSome := false
	seedTrytes, _ := gadk.ToTrytes(*seed)
	candidates := make(map[string]int)

	for index := 0; index < countAddresses; index++ {
		adr, _ := gadk.NewAddress(seedTrytes, index, 2)
		if addressToValueMap[string(adr)] > 0 {
			candidates[string(adr)] = index
			GUILog(string(adr) + " MIGRATION CANDIDATE (had balance at genesis)")
			foundSome = true
		} else {
		}
		if index%50 == 49 {
			GUILog("Scanned", index+1, "of", countAddresses, "addresses")
		}
	}

	zero := big.NewInt(0)
	//use fixed timestamp so each bundle is fixed and reduces 'reuse' issue.
	layout := "2006-01-02T15:04:05.000Z"
	str := "2021-11-11T00:00:00.000Z"
	t, _ := time.Parse(layout, str)
	msgshown := false
	someError := false

	if len(candidates) > 0 {
		GUILog("Found ADK migration candidates:", len(candidates), " addresses. Now checking migration status.")
		//
		vAGSClaimContract := common.HexToAddress("0xdf61877BaB8B3789A41924fdDB01804cfAED3C55")
		vTokenContract := common.HexToAddress("0xB56a995117b41dA5BA90DA5c66116c60c6e19170")
		cAGSContract, _ := NewAGSClaim(vAGSClaimContract, client)
		cADKToken, _ := NewADKToken(vTokenContract, client)

		// validate address first
		az9Addr, _ := cADKToken.ADDRTOAZ9(nil, common.HexToAddress(*targetAirdrop))
		az9Addr_check, _ := cADKToken.AZ9TOADDR(nil, az9Addr)

		az9AddrTrytesA, _ := gadk.ToAddress(az9Addr)

		if strings.ToLower(az9Addr_check.Hex()) != strings.ToLower(*targetAirdrop) {
			GUILog("ERROR: Target claim address error. Invalid: ", *targetAirdrop, az9Addr_check)
			return successfullyClaimed, false
		}

		api := gadk.NewAPI("http://api1.mainnet.aidoskuneen.com:14265", nil) // dummy, not really used
		bulkProcessing := false

		big_1000000000000000000 := big.NewInt(10)
		big_1000000000000000000.Exp(big_1000000000000000000, big.NewInt(18), nil)

		for adr, addrindex := range candidates {
			claimable := big.NewInt(0)
			claimable, _ = cAGSContract.ClaimableAZ9(nil, adr)
			adrTrytesA, _ := gadk.ToAddress(adr)
			if claimable.Cmp(zero) > 0 && !msgshown {
				msgshown = true
				GUILog("**************************************************")
				GUILog("** GREAT! You have old ADK that can be claimed.  *")
				GUILog("**************************************************")

			}
			GUILog(adr, "claimable:", claimable, "uADK (~", big.NewInt(1).Div(claimable, big_1000000000000000000), "ADK)")

			if claimable.Cmp(zero) > 0 {
				GUILog("Generating claim transaction, sending to", *targetAirdrop, "\n      =", az9Addr)
				adKey := gadk.NewKey(seedTrytes, addrindex, 2)
				trs := []gadk.Transfer{
					gadk.Transfer{
						Address: az9AddrTrytesA,
						Value:   0,
						Message: "CLAIMTRANSACTION",
					},
					gadk.Transfer{
						Address: adrTrytesA,
						Value:   0,
						Message: "X",
					},
					gadk.Transfer{
						Address: adrTrytesA,
						Value:   0,
						Message: "X",
					},
				}
				bundle, _ := gadk.PrepareTransfers(api, seedTrytes, trs, nil, az9AddrTrytesA, 2)
				bundle[1].Timestamp = t
				bundle[2].Timestamp = t
				bundle.Finalize([]gadk.Trytes{bundle[0].Trytes(), bundle[0].Trytes(), bundle[0].Trytes()})
				bundleHash := bundle.Hash()
				nHash := bundle.Hash().Normalize()
				bundle[1].SignatureMessageFragment = gadk.Sign(nHash[:27], adKey[:6561/3])
				bundle[2].SignatureMessageFragment = gadk.Sign(nHash[27:2*27], adKey[6561/3:2*6561/3])
				//
				if gadk.IsValidSig(adrTrytesA, []gadk.Trytes{bundle[1].SignatureMessageFragment, bundle[2].SignatureMessageFragment}, bundleHash) {
					GUILog("Key Valid")
				} else {
					GUILog("Key Invalid")
					return successfullyClaimed, false
				}

				GUILog(bundle.Hash())
				_, pow := gadk.GetBestPoW()
				GUILog("PoW (1/3)")
				bundle[0].Nonce, _ = pow(bundle[0].Trytes(), 15)
				GUILog("PoW (2/3)")
				bundle[1].Nonce, _ = pow(bundle[1].Trytes(), 15)
				GUILog("PoW (3/3)")
				bundle[2].Nonce, _ = pow(bundle[2].Trytes(), 15)
				t1 := bundle[0].Trytes()
				t2 := bundle[1].Trytes()
				t3 := bundle[2].Trytes()
				//GUILog(t1,t2,t3)
				//
				GUILog("Claiming...")
				t_opt := GetAuthClaim(client)

				tx, errX := cAGSContract.PostTransactions(t_opt, string(t1)+string(t2)+string(t3))
				if errX != nil {
					GUILog(err)
					//return successfullyClaimed
					someError = true
					continue
				}
				GUILog("tx sent: " + tx.Hash().Hex())
				if !bulkProcessing {
					receipt, err := bind.WaitMined(context.Background(), client, tx)
					if err != nil {
						GUILog(err)
						someError = true
						//return successfullyClaimed
						continue
					}
					if receipt.Status != 1 { // reverted, extract the revert reason
						GUILog("TX Error : ")
						GUILog("tx.From() : ", t_opt.From)
						GUILog("tx.To() : ", tx.To())
						GUILog("tx.Gas() : ", tx.Gas())
						GUILog("tx.GasPrice() : ", tx.GasPrice())
						GUILog("tx.Value() : ", tx.Value())
						GUILog("tx.Data() : ", tx.Data())
						GUILog("data : ", string(t1)+string(t2)+string(t3))
						msg := ethereum.CallMsg{
							From:     t_opt.From,
							To:       tx.To(),
							Gas:      tx.Gas(),
							GasPrice: tx.GasPrice(),
							Value:    tx.Value(),
							Data:     tx.Data(),
						}
						result, err := client.CallContract(context.Background(), msg, receipt.BlockNumber)
						GUILog("Result:", result)
						GUILog("ResultErr:", err)
						someError = true
					} else {
						successfullyClaimed.Add(successfullyClaimed, tx.Value())
						GUILog("TRANSACTION MINED")
					}
				}
			}
		}
	}

	if !foundSome {
		GUILog("No claimable ADK found.")
	}

	GUILog("Scan/migration completed.")

	time.Sleep(5 * time.Second)
	return successfullyClaimed, !someError
}
