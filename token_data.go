package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/txaty/go-merkletree"
)

const (
	AuraNone        = iota // 0
	AuraLow                // 1
	AuraMed                // 2
	AuraHigh               // 3
	AuraOverflowing        // 4
)

const (
	FukushimaFishIPFSHash = "QmVzguHhExH1YKEhf4jUJzFX3zE1McXaKQU1CPS4Ntd3yP"

	TraitAura       = "Aura"
	TraitBackground = "Background"
	TraitCelestial  = "Celestials"

	FlagCelestial = 0x01
)

var (
	AuraMapping = map[string]int{
		"none":        AuraNone,
		"low":         AuraLow,
		"medium":      AuraMed,
		"high":        AuraHigh,
		"overflowing": AuraOverflowing,
	}

	LevelToAuraMapping = map[int]string{
		AuraNone:        "none",
		AuraLow:         "low",
		AuraMed:         "medium",
		AuraHigh:        "high",
		AuraOverflowing: "overflowing",
	}

	BackgroundModifier = 0x01
)

type FishTokenData struct {
	TokenID    int      `json:"tokenId"`
	AuraLevel  int      `json:"auraLevel"`
	Background string   `json:"background"`
	Celestial  bool     `json:"celestial"`
	Proof      []string `json:"proof"`
	Path       int      `json:"path"`
	Encoded    string   `json:"encodedData"`
}

func (d *FishTokenData) Serialize() ([]byte, error) {

	// We're going to be packing
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	encodeMe := abi.Arguments{
		abi.Argument{
			Name: "uint256",
			Type: uint256Type,
		},
	}

	/*
	   uint256 encoded = uint256(uint16(flags));

	   encoded <<= 16;
	   encoded |= level;

	   encoded <<= 16;
	   encoded |= token;
	*/

	var flags uint16

	if d.Celestial {
		flags |= FlagCelestial
	}

	// THINGS WE ARE STORING: Token Id
	metadata := big.NewInt(int64(flags))
	// <<= 16
	metadata = new(big.Int).Lsh(metadata, 16)
	// |= level
	metadata = new(big.Int).Or(metadata, big.NewInt(int64(d.AuraLevel)))

	// <<= 16
	metadata = new(big.Int).Lsh(metadata, 16)
	// |= token
	metadata = new(big.Int).Or(metadata, big.NewInt(int64(d.TokenID)))

	// return abi.encode(data)
	return encodeMe.Pack(metadata)
}

func generateTokenData(tokenId int) *FishTokenData {
	metadata, err := fetchMetadata(FukushimaFishIPFSHash, tokenId)

	if err != nil {
		log.Fatalln(err)
	}

	data := &FishTokenData{
		TokenID: tokenId,
	}

	for _, attr := range metadata.Attributes {
		if attr.TraitType == TraitAura {
			level, ok := AuraMapping[strings.ToLower(attr.Value)]
			if !ok {
				log.Fatalln("unknown trait type for aura found:", attr.Value)
			}
			data.AuraLevel = level
		} else if attr.TraitType == TraitBackground {
			data.Background = attr.Value
		} else if attr.TraitType == TraitCelestial {
			data.Celestial = true
			// give celestials Aura High for now
			data.AuraLevel = AuraHigh
		}
	}

	return data
}

func generateTokenDataMerkleTree() {

	log.Println(LevelToAuraMapping)

	log.Println("fetching IPFS data.")

	tokens := []merkletree.DataBlock{}
	for i := 1; i <= 3888; i++ {
		tokenData := generateTokenData(i)
		tokens = append(tokens, tokenData)

		fmt.Print("\rfetched information for token: ", i, "\r")
	}

	log.Println("generating merkle tree")
	merkleTree, err := merkletree.New(&merkletree.Config{
		HashFunc: func(b []byte) ([]byte, error) {
			return crypto.Keccak256(b), nil
		},
		Mode:          merkletree.ModeProofGenAndTreeBuild,
		RunInParallel: false,
		NoDuplicates:  false,
	}, tokens)

	if err != nil {
		log.Fatalln(err)
	}

	for _, token := range tokens {
		proof, err := merkleTree.GenerateProof(token)

		if err != nil {
			log.Fatalln(err)
		}

		var proofHex []string

		for _, proof := range proof.Siblings {
			proofHex = append(proofHex, hexutil.Encode(proof))
		}

		token.(*FishTokenData).Proof = proofHex
		token.(*FishTokenData).Path = int(proof.Path)

		encoded, _ := token.Serialize()
		token.(*FishTokenData).Encoded = hexutil.Encode(encoded)
	}

	encoded := map[string]interface{}{}

	encoded["root"] = hexutil.Encode(merkleTree.Root)
	encoded["data"] = tokens

	jsonEncoded, _ := json.MarshalIndent(encoded, " ", "    ")
	os.WriteFile("fish_token_data.json", jsonEncoded, os.ModePerm)
}
