package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjaswtf/fukushimafish/erc721"
	"github.com/txaty/go-merkletree"
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
	FlagWhitelist  = flag.Bool("whitelist", false, "generates the whitelist files if specified")
	FlagOracle     = flag.Bool("oracle", false, "feeds the data contract the necessary data")
)

type Snapshot struct {
	// This is the merkle root hash of the whitelisted addresses
	MerkleRootHash string `json:"rootHash"`
	// these are the unique owners of the NFT
	UniqueAddresses uint `json:"uniqueAddresses"`
	// Number of whitelisted addresses
	TotalWhitelistedAddresses uint `json:"totalWhitelistedAddresses"`
	// this is a collection of information pulled from on-chain & IPFS for a given address
	Snapshot []SnapshotData `json:"data"`
}

type SnapshotData struct {
	Address          string   `json:"address"` // hex encoded address
	NumberOfNFTOwned int      `json:"numberOwned"`
	HasKOI           bool     `json:"hasKoi"`
	MerkleTreeIndex  uint     `json:"path,omitempty"`
	Proof            []string `json:"proof,omitempty"` // hex encoded proofs for the whitelist
}

func (d *SnapshotData) Serialize() ([]byte, error) {
	return common.HexToAddress(d.Address).Bytes(), nil
}

func (snapshotData *SnapshotData) HasWhitelist() bool {
	return snapshotData.HasKOI || snapshotData.NumberOfNFTOwned >= 5
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

func generateWhitelist() {

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
	snapshot := map[common.Address]*SnapshotData{}

	uniqueAddresses := 0

	for i := 0; i < int(totalSupply.Uint64()); i++ {
		token := i + 1
		owner, _ := contract.OwnerOf(nil, big.NewInt(int64(token)))
		snapshotData, ok := snapshot[owner]

		if !ok {
			snapshotData = &SnapshotData{
				Address:          owner.Hex(),
				NumberOfNFTOwned: 0,
			}
			snapshot[owner] = snapshotData
			uniqueAddresses++
		}

		snapshotData.NumberOfNFTOwned += 1

		metadata, err := fetchMetadata(BearRidersNFTIPFSHash, token)

		if err != nil {
			log.Fatalln(token, err)
		}

		snapshotData.HasKOI = hasKoiCoin(metadata)
		fmt.Print("\rfetched information for token: ", token, "\r")
	}

	log.Println("generating leaves for the merkle tree")

	// push whitelisted addresses to this list
	var leaves []merkletree.DataBlock

	// in the case of multiple addresses, we wouldn't want to add those twice
	// keep track of what addresses have been pushed to the leaves
	whitelistedAddresses := map[string]bool{}

	for _, snapshotData := range snapshot {
		if snapshotData.HasWhitelist() && !whitelistedAddresses[snapshotData.Address] {
			whitelistedAddresses[snapshotData.Address] = true
			leaves = append(leaves, snapshotData)
		}
	}

	merkleTree, _ := merkletree.New(&merkletree.Config{
		HashFunc: func(b []byte) ([]byte, error) {
			return crypto.Keccak256(b), nil
		},
		Mode:          merkletree.ModeProofGenAndTreeBuild,
		RunInParallel: false,
		NoDuplicates:  false,
	}, leaves)

	var finalSnapshot Snapshot

	finalSnapshot.UniqueAddresses = uint(uniqueAddresses)

	finalSnapshot.TotalWhitelistedAddresses = uint(len(leaves))

	finalSnapshot.MerkleRootHash = hexutil.Encode(merkleTree.Root)

	for _, snapshotData := range snapshot {
		if snapshotData.HasWhitelist() {

			generatedProof, err := merkleTree.GenerateProof(snapshotData)
			if err != nil {
				log.Fatalln("failed to create proof for", snapshotData.Address)
			}
			var proofHex []string

			for _, proof := range generatedProof.Siblings {
				proofHex = append(proofHex, hexutil.Encode(proof))
			}

			snapshotData.MerkleTreeIndex = uint(generatedProof.Path)
			snapshotData.Proof = proofHex
		}
		finalSnapshot.Snapshot = append(finalSnapshot.Snapshot, *snapshotData)
	}

	encoded, _ := json.Marshal(finalSnapshot)

	ioutil.WriteFile(*FlagOutputFile, encoded, os.ModePerm)
}

func main() {

	rand.Seed(time.Now().UnixMilli())
	// push whitelisted addresses to this list

	var levelsTree []merkletree.DataBlock

	// this would iterate through every single token
	for i := 0; i < 10; i++ {

		// here we would parse the metadata
		// metadata := fetchMetadata(fukushimaFishIPFSHash, token)

		// radiationLevel := metadata.Attributes.Find("radiationLevel")
		radiationLevel := rand.Intn(6)

		data := &FukushimaFishData{
			TokenID:        i,
			RadiationLevel: radiationLevel * 2,
		}

		// bg := metadata.Attributes.Find("background")
		// if bg == "reactor" {

		if rand.Int()%10 == 0 {
			data.RadiationLevel += 1 // add the reactor modifier
		}

		log.Println("generated token:", i, "with radiation level of", data.RadiationLevel)

		levelsTree = append(levelsTree, data)
	}

	merkleTree, _ := merkletree.New(&merkletree.Config{
		HashFunc: func(b []byte) ([]byte, error) {
			return crypto.Keccak256(b), nil
		},
		Mode:          merkletree.ModeProofGenAndTreeBuild,
		RunInParallel: false,
		NoDuplicates:  false,
	}, levelsTree)

	for token, dataBlock := range levelsTree {
		proof, _ := merkleTree.GenerateProof(dataBlock)
		var proofHex []string

		for _, proof := range proof.Siblings {
			proofHex = append(proofHex, hexutil.Encode(proof))
		}

		log.Println(merkleTree.Verify(dataBlock, proof))
		log.Println("proof for token:", token, "path:", proof.Path, "proof:", proofHex)
	}

	log.Println("root hash:", hexutil.Encode(merkleTree.Root))

}
