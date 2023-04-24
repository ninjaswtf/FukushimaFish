package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjaswtf/fukushimafish/erc721"
	"github.com/txaty/go-merkletree"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	BearRidersNFTContract   = "0x193fcbc5Cd9B28B75e54533647db9d61F6f66812"
	BearRidersNFTIPFSHash   = "Qmd8bo3viapTezspL9a1wSWKNiwZcriK1hpcaP4q4xZE5F"
	BearRidersNFTIPFSHashV2 = "bafybeig3zee5snddawt2aikacqihiucn2kw62ibvviyzvuuu4j3y5kn7ti"
)

type RadiationLevels int

const (
	NONE RadiationLevels = iota // default value
	LOW
	MED
	HIGH
	DANGER
	REACTOR
)

var (
	IPFSMetadataEndpoints = []string{
		// it is best if you run your own IPFS node.
		"http://127.0.0.1:8080/ipfs/%s/%d.json",
		"https://gateway.pinata.cloud/ipfs/%s/%d.json",
		"https://gateway.ipfs.io/ipfs/%s/%d.json",
		"https://cloudflare-ipfs.com/ipfs/%s/%d.json",
	}

	Database *gorm.DB
)

func fetchMetadata(ipfsHash string, token int) (*Metadata, error) {
	for i := 0; i < len(IPFSMetadataEndpoints); i++ {
		resp, err := http.Get(fmt.Sprintf(IPFSMetadataEndpoints[i], ipfsHash, token))

		if resp.StatusCode != 200 {
			time.Sleep(500)
			log.Println("failed to fetch, retrying with another endpoint")
			continue
		}
		if err != nil {
			return nil, err
		}

		bytesBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var metadata Metadata

		if err = json.Unmarshal(bytesBody, &metadata); err != nil {
			return nil, err
		}

		return &metadata, nil
	}

	return nil, errors.New("failed to fetch metadata")
}

func hasKoiCoin(metadata *Metadata) bool {
	const (
		CoinsFemale = "Coins Female"
		CoinsMale   = "Coins Male"

		KoiCoinFemale = "BR_Coin_KOI_F"
		KoiCoinMale   = "BR_Coin_KOI"
	)
	for _, attr := range metadata.Attributes {
		if attr.TraitType == CoinsFemale || attr.TraitType == CoinsMale {
			return attr.Value == KoiCoinFemale || attr.Value == KoiCoinMale
		}
	}
	return false
}

var (
	FlagOutputFile = flag.String("output", "snapshot.json", "set the output file")
	FlagWhitelist  = flag.Bool("generate-whitelist", false, "generates the whitelist files if specified")
	FlagSnapshot   = flag.Bool("snapshot", false, "takes a snapshot of the BearRiders NFT")
	FlagOracle     = flag.Bool("oracle", false, "feeds the data contract the necessary data")

	FlagWhitelistAddress = flag.Bool("whitelist", false, "adds the specified address to the whitelist")
	FlagAddress          = flag.String("address", "", "0x EVM address")
	FlagWhitelistSlots   = flag.Int("slots", 0, "number of whitelist slots")
)

type Snapshot struct {
	// This is the merkle root hash of the whitelisted addresses
	MerkleRootHash string `json:"rootHash"`
	// these are the unique owners of the NFT
	UniqueAddresses uint `json:"uniqueAddresses"`
	// Number of whitelisted addresses
	TotalWhitelistedAddresses uint `json:"totalWhitelistedAddresses"`
	// this is a collection of information pulled from on-chain & IPFS for a given address
	Snapshot []*WhitelistData `json:"data"`
}

type WhitelistData struct {
	Address         string      `json:"address" gorm:"primaryKey;column:address;"` // hex encoded address
	WhitelistSlots  int         `json:"whitelistSlots" gorm:"column:whitelist_slots;"`
	OwnedTokens     int         `json:"ownedTokens" gorm:"column:owned_tokens;"`
	HasKOI          bool        `json:"hasKoi" gorm:"column:has_koi;"`
	MerkleTreeIndex uint        `json:"path,omitempty" gorm:"column:merkle_tree_index;"`
	Proof           StringArray `json:"proof,omitempty" gorm:"column:proof;"` // hex encoded proofs for the whitelist
}

func (*WhitelistData) TableName() string {
	return "whitelist"
}

func (d *WhitelistData) Serialize() ([]byte, error) {
	addressType, _ := abi.NewType("address", "address", nil)
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	encodeMe := abi.Arguments{
		abi.Argument{
			Name: "address",
			Type: addressType,
		},
		abi.Argument{
			Name: "uint256",
			Type: uint256Type,
		},
	}

	encoded, _ := encodeMe.Pack(common.HexToAddress(d.Address), big.NewInt(int64(d.WhitelistSlots)))

	log.Println("encoding:", encoded, "address:", d.Address)

	return encodeMe.Pack(common.HexToAddress(d.Address), big.NewInt(int64(d.WhitelistSlots)))
}

func (snapshotData *WhitelistData) HasWhitelist() bool {
	return snapshotData.HasKOI || snapshotData.WhitelistSlots >= 1
}

type FukushimaFishData struct {
	TokenID        int
	RadiationLevel int
}

func (d *FukushimaFishData) Serialize() ([]byte, error) {
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	encodeMe := abi.Arguments{
		abi.Argument{
			Name: "tokenId",
			Type: uint256Type,
		},
		abi.Argument{
			Name: "radiationLevel",
			Type: uint256Type,
		},
	}
	return encodeMe.Pack(big.NewInt(int64(d.TokenID)), big.NewInt(int64(d.RadiationLevel)))
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Metadata struct {
	Attributes []Attribute `json:"attributes"`
}

func init() {
	flag.Parse()

	db, err := gorm.Open(sqlite.Open("whitelist.db"))

	if err != nil {
		log.Fatalln(err)
	}

	Database = db

	Database.AutoMigrate(&WhitelistData{})
}

func takeSnapshot() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/c6b15721b1044ab7a30d3b911f535e47")

	if err != nil {
		log.Fatalln(err)
	}

	contract, err := erc721.NewERC721(common.HexToAddress(BearRidersNFTContract), client)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("fetching on-chain & IPFS data.")
	totalSupply, _ := contract.TotalSupply(nil)
	snapshot := map[common.Address]*WhitelistData{}

	uniqueAddresses := 0

	for i := 0; i < int(totalSupply.Uint64()); i++ {
		token := i + 1
		owner, _ := contract.OwnerOf(nil, big.NewInt(int64(token)))
		snapshotData, ok := snapshot[owner]

		if !ok {
			snapshotData = &WhitelistData{
				Address:        owner.Hex(),
				WhitelistSlots: 0,
			}
			snapshot[owner] = snapshotData
			uniqueAddresses++
		}

		snapshotData.OwnedTokens += 1

		metadata, err := fetchMetadata(BearRidersNFTIPFSHash, token)

		if err != nil {
			log.Fatalln(token, err)
		}

		snapshotData.HasKOI = hasKoiCoin(metadata)

		if snapshotData.HasKOI {
			snapshotData.WhitelistSlots += 1
		}

		Database.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "address"}},
			DoUpdates: clause.AssignmentColumns([]string{"whitelist_slots", "has_koi", "owned_tokens"}),
		}).Create(snapshotData)

		fmt.Print("\rfetched information for token: ", token, "\r")
	}

	// for _, data := range snapshot {
	// 	Database.Create(data)
	// }

}

func generateWhitelist() {
	log.Println("generating leaves for the merkle tree")

	var snapshot []*WhitelistData
	Database.Find(&snapshot)

	// push whitelisted addresses to this list
	var leaves []merkletree.DataBlock

	// in the case of multiple addresses, we wouldn't want to add those twice
	// keep track of what addresses have been pushed to the leaves
	whitelistedAddresses := map[string]bool{}

	for _, snapshotData := range snapshot {

		if snapshotData.HasWhitelist() && !whitelistedAddresses[snapshotData.Address] {

			if snapshotData.Address == "0xF42D1c0c0165AF5625b2ecD5027c5C5554e5b039" {
				log.Println("expected address has whitelist")
			}
			whitelistedAddresses[snapshotData.Address] = true
			leaves = append(leaves, snapshotData)
		}
	}

	merkleTree, err := merkletree.New(&merkletree.Config{
		HashFunc: func(b []byte) ([]byte, error) {
			return crypto.Keccak256(b), nil
		},
		Mode:          merkletree.ModeProofGenAndTreeBuild,
		RunInParallel: false,
		NoDuplicates:  false,
	}, leaves)

	if err != nil {
		log.Fatalln(err)
	}

	var finalSnapshot Snapshot

	// TODO: SQL query for calculating the Unique Addresses
	// finalSnapshot.UniqueAddresses = uint(uniqueAddresses)

	finalSnapshot.TotalWhitelistedAddresses = uint(len(leaves))

	finalSnapshot.MerkleRootHash = hexutil.Encode(merkleTree.Root)

	log.Println(merkleTree.Root)

	for _, snapshotData := range snapshot {
		if snapshotData.HasWhitelist() {

			log.Println(&snapshotData)

			generatedProof, err := merkleTree.GenerateProof(snapshotData)
			if err != nil {
				log.Fatalln("failed to create proof for", snapshotData.Address, ":", err)
			}
			var proofHex []string

			for _, proof := range generatedProof.Siblings {
				proofHex = append(proofHex, hexutil.Encode(proof))
			}

			snapshotData.MerkleTreeIndex = uint(generatedProof.Path)
			snapshotData.Proof = proofHex
		}
		finalSnapshot.Snapshot = append(finalSnapshot.Snapshot, snapshotData)
	}

	encoded, _ := json.Marshal(finalSnapshot)

	ioutil.WriteFile(*FlagOutputFile, encoded, os.ModePerm)
}

func main() {

	done := make(chan struct{})

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		close(done)
	}()

	go func() {
		if *FlagWhitelistAddress {

			if *FlagAddress == "" || *FlagWhitelistSlots <= 0 {
				log.Fatalln("invalid parameters")
			}

			addr := common.HexToAddress(*FlagAddress)
			Database.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "address"}},
				DoUpdates: clause.AssignmentColumns([]string{"whitelist_slots"}),
			}).Create(&WhitelistData{
				Address:        addr.String(),
				WhitelistSlots: *FlagWhitelistSlots,
			})

			log.Printf("address %s has been given %d whitelist slot(s)\n", addr.String(), *FlagWhitelistSlots)

			done <- struct{}{}
			return
		}

		if *FlagWhitelist {
			generateWhitelist()
			done <- struct{}{}
			return
		}

		if *FlagSnapshot {
			takeSnapshot()

			done <- struct{}{}
			return
		}
	}()

	<-done
}
