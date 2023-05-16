package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ninjaswtf/fukushimafish/fish"
)

var (
	KeyPathFlag     = flag.String("key", "key.json", "the key for the deployer account")
	KeyPasswordFlag = flag.String("password", "password", "the password to unlock the key")
	CreateKeyFlag   = flag.Bool("create-key", false, "creates a new private key at the deployer key path")
)

func CreateKey(password string) (string, common.Address, error) {

	ks := keystore.NewKeyStore("keys", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)

	if err != nil {
		return "", common.Address{}, err
	}

	jsonKey, err := ks.Export(account, password, password)

	if err != nil {
		return "", common.Address{}, err
	}

	return string(jsonKey), account.Address, err
}

func initKey() *keystore.Key {
	jsonKey, err := ioutil.ReadFile(*KeyPathFlag)

	if err != nil {
		log.Fatalln("could not find the key file", err)
	}

	key, err := keystore.DecryptKey(jsonKey, *KeyPasswordFlag)

	if err != nil {
		log.Fatalln("failed to decrypt key:", err)
	}

	log.Printf("using account %s, make sure this is well-funded!\n", key.Address.String())
	return key
}

func airdropSnapshot() {

	mainnetClient, _ := ethclient.Dial("https://mainnet.infura.io/v3/c6b15721b1044ab7a30d3b911f535e47")
	oldContract, _ := fish.NewFish(common.HexToAddress("0x5667B16275eFc836B5e3298409cc9c949fA38970"), mainnetClient)

	// address => ownedTokens
	ownedTokens := map[string][]int{}

	for i := 1; i <= 323; i++ {
		owner, _ := oldContract.OwnerOf(nil, big.NewInt(int64(i)))

		ownedTokensArray := ownedTokens[owner.Hex()]

		ownedTokens[owner.Hex()] = append(ownedTokensArray, i)

		fmt.Print("\rfetched information for token: ", i, "\r")
	}

	encoded, _ := json.MarshalIndent(ownedTokens, " ", "    ")

	os.WriteFile("airdrop_snapshot.json", encoded, os.ModePerm)
}

func mainnetEthGasCost(gasUsedInt uint64, mainnetClient *ethclient.Client) *big.Float {
	gasUsedString := strconv.Itoa(int(gasUsedInt))
	gasUsed, _ := new(big.Int).SetString(gasUsedString, 10)
	gasCost, _ := mainnetClient.SuggestGasPrice(context.Background())

	gasPriceInEth := new(big.Int).Mul(gasUsed, gasCost)

	gasPriceFloat, _ := new(big.Float).SetString(gasPriceInEth.String())
	etherFloat := big.NewFloat(params.Ether)

	etherResult := new(big.Float).Quo(gasPriceFloat, etherFloat)

	return etherResult
}

type StringCount struct {
	Value string `json:"address"`
	Count int    `json:"amount"`
}

// brain still tired, asked chat GPT to write this lol
func countConsecutiveStrings(arr []string) []StringCount {
	if len(arr) == 0 {
		return nil
	}

	counts := []StringCount{}
	currentCount := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			currentCount++
		} else {
			counts = append(counts, StringCount{arr[i-1], currentCount})
			currentCount = 1
		}
	}

	counts = append(counts, StringCount{arr[len(arr)-1], currentCount})
	return counts
}

func airdrop() {

	ownedTokens := map[string][]int{}

	data, _ := os.ReadFile("airdrop_snapshot.json")

	json.Unmarshal(data, &ownedTokens)

	// key := initKey()

	// mainnetClient, _ := ethclient.Dial("https://mainnet.infura.io/v3/c6b15721b1044ab7a30d3b911f535e47")
	// testnetClient, _ := ethclient.Dial("https://sepolia.infura.io/v3/c6b15721b1044ab7a30d3b911f535e47")

	// nftContract := common.HexToAddress("0xEe259D1aae364bCb2D28333c712d394F052984c3")
	// oldContract, _ := fish.NewFish(common.HexToAddress("0x5667B16275eFc836B5e3298409cc9c949fA38970"), mainnetClient)
	// newContract, _ := fish.NewFish(nftContract, mainnetClient)

	// transactionOps, _ := bind.NewKeyedTransactorWithChainID(key.PrivateKey, big.NewInt(1))

	/*
			Neo	0x5e2636ed16f900503f1acfbf595c7e739b67a2b3	30
		Atomik	0xa94b9D00E62B3e1d61FFa8a40338D8D63a604A8B	15
		Ben	0xB895D84cCE113B1889Dc7554bcB112701577bB40	18
		Fauve	0x7d233c0F9fc98BbABDf20FdD3619B51cf3816886	10
		Ryan	0x180161e9b47faea68f7d4d7e618837416ab39a64	7
		0xBill	0xC9D4cD1144341Ff8156a41472B4e65261724710B	6
		0xMissy	0xc634597fcf2a7bfb4bcb1e1761b3761bfdbcd89c	6
		Seize	0x43Ec5640E18F7384761d8817AA55d38C9a03D855	6
		Takuya	0x1afc21d5086f03826cc89082e74746ee463690ab	6
		Celestial Studios	0x5eCb3dBE838e647524F0Cb46F2027A44Cdda3021	65
		NFTCanvases	0xb44Df61BA4B5Dd10869F75232667c9a0F3691211	5
		Birthmark	0xFD2543F11B6AAB799C5e55f977bC24B54bf16C21	5
		Cyan	0xD66FE0a25b5Ef9E916E1A035F00fE7601838182C	5
		Mistori	0x3426ffa3dD023Ff9b361D7E5666dCf77794bcC71	5
		Side	Send To Celestial	2
		8020	Send To Celestial	4
		STYLE	0x1Fce506e29f6422d769eC9d5442AE1356742847e	5
	*/

	type mint struct {
		Address string `json:"address"`
		Amount  int    `json:"amount"`
	}
	var orderedMints []mint = []mint{
		// {"0x5e2636ed16f900503f1acfbf595c7e739b67a2b3", 30}, // neo
		// {"0xa94b9D00E62B3e1d61FFa8a40338D8D63a604A8B", 15}, // atomik
		// {"0xB895D84cCE113B1889Dc7554bcB112701577bB40", 18}, // ben
		// {"0x7d233c0F9fc98BbABDf20FdD3619B51cf3816886", 10}, // fauve
		// {"0xF42D1c0c0165AF5625b2ecD5027c5C5554e5b039", 7},  // Ryan
		// {"0xC9D4cD1144341Ff8156a41472B4e65261724710B", 6},  // 0xBill
		// {"0xc634597fcf2a7bfb4bcb1e1761b3761bfdbcd89c", 6},  // 0xMissy
		// {"0x43Ec5640E18F7384761d8817AA55d38C9a03D855", 6},  // sieze
		// {"0x34319d7f5C8fB67d37566BD5bDC3f5445590cdC6", 6},  // Takuya
		// {"0x5eCb3dBE838e647524F0Cb46F2027A44Cdda3021", 65}, // Celestial Studios
		// {"0xb44Df61BA4B5Dd10869F75232667c9a0F3691211", 5},  // NFT Canvases
		// {"0xFD2543F11B6AAB799C5e55f977bC24B54bf16C21", 5},  // Birthmark
		// {"0xD66FE0a25b5Ef9E916E1A035F00fE7601838182C", 5},  // Cyan
		// {"0x3426ffa3dD023Ff9b361D7E5666dCf77794bcC71", 5},  // Mistori
		// {"0x5eCb3dBE838e647524F0Cb46F2027A44Cdda3021", 6},  // Send To Celestial (6)
		// {"0x1Fce506e29f6422d769eC9d5442AE1356742847e", 5},  // STYLE
	}

	var teamWallets map[string]bool = map[string]bool{
		"0x5e2636ed16f900503f1acfbf595c7e739b67a2b3": true,
		"0xa94b9d00e62b3e1d61ffa8a40338d8d63a604a8b": true,
		"0xb895d84cce113b1889dc7554bcb112701577bb40": true,
		"0x7d233c0f9fc98bbabdf20fdd3619b51cf3816886": true,
		"0xf42d1c0c0165af5625b2ecd5027c5c5554e5b039": true,
		"0xc9d4cd1144341ff8156a41472b4e65261724710b": true,
		"0xc634597fcf2a7bfb4bcb1e1761b3761bfdbcd89c": true,
		"0x43ec5640e18f7384761d8817aa55d38c9a03d855": true,
		"0x34319d7f5c8fb67d37566bd5bdc3f5445590cdc6": true,
		"0x5ecb3dbe838e647524f0cb46f2027a44cdda3021": true,
		"0xb44df61ba4b5dd10869f75232667c9a0f3691211": true,
		"0xfd2543f11b6aab799c5e55f977bc24b54bf16c21": true,
		"0xd66fe0a25b5ef9e916e1a035f00fe7601838182c": true,
		"0x3426ffa3dd023ff9b361d7e5666dcf77794bcc71": true,
		"0x1fce506e29f6422d769ec9d5442ae1356742847e": true,
	}

	// token id => address

	totalTokens := 0

	for address, tokens := range ownedTokens {
		if teamWallets[strings.ToLower(address)] {
			continue
		}
		orderedMints = append(orderedMints, mint{address, len(tokens)})

		totalTokens += len(tokens)
	}

	log.Println(totalTokens)
	log.Println(len(orderedMints) == len(ownedTokens)-15, len(orderedMints), len(ownedTokens))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(orderedMints), func(i, j int) {
		orderedMints[i], orderedMints[j] = orderedMints[j], orderedMints[i]
	})

	// log.Println(len(orderedMints) == len(ownedTokens)-15, len(orderedMints), len(ownedTokens))

	// encoded, _ := json.MarshalIndent(countConsecutiveStrings(sortedAddresses), " ", "    ")

	// os.WriteFile("ordered_mints.json", encoded, os.ModePerm)

	// totalGasUsedInEther := new(big.Float)
	// // Mints transactions in order they were originally minted in
	// for _, mint := range orderedMints {

	// 	log.Println("minting to address:", mint.Address)

	// 	txn, _ := newContract.MintAirdropTeamMint(transactionOps, common.HexToAddress(mint.Address),
	// 		big.NewInt(int64(mint.Amount)))

	// 	receipt, err := bind.WaitMined(context.Background(), mainnetClient, txn)

	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	ethGasCost := mainnetEthGasCost(receipt.GasUsed, mainnetClient)
	// 	totalGasUsedInEther = new(big.Float).Add(totalGasUsedInEther, ethGasCost)
	// 	log.Println("minted", mint.Amount, "nfts to team member", mint.Address, "for eth cost:\t", ethGasCost.String())

	// 	log.Println("total eth gas cost so far:", totalGasUsedInEther.String())

	// 	log.Println("waiting for 30 seconds to cool down")
	// 	time.Sleep(30 * time.Second)
	// 	log.Println("")
	// }
}
