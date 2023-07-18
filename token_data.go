package main

import (
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	FukushimaFishIPFSHash = "QmVzguHhExH1YKEhf4jUJzFX3zE1McXaKQU1CPS4Ntd3yP"

	AuraNone        = iota * 2 // 0
	AuraLow                    // 2
	AuraMed                    // 4
	AuraHigh                   // 6
	AuraOverflowing            // 8

	TraitAura       = "Aura"
	TraitBackground = "Background"
	TraitCelestial  = "Celestials"
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
	TokenID    int
	AuraLevel  int
	Background string
	Celestial  bool
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

	// THINGS WE ARE STORING: Token Id
	nInt := big.NewInt(int64(d.TokenID))

	// return encodeMe.Pack(common.HexToAddress(d.Address), big.NewInt(int64(d.WhitelistSlots)))
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
