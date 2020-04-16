package auscash

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	// MainnetMagic is mainnet network constant
	MainnetMagic wire.BitcoinNet = 0xcbf2c6f1
	// TestnetMagic is testnet network constant
	TestnetMagic wire.BitcoinNet = 0xf1c8d2fd
	// RegtestMagic is regtest network constant
	RegtestMagic wire.BitcoinNet = 0xdab5bffa
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{23}
	MainNetParams.ScriptHashAddrID = []byte{63}
	MainNetParams.Bech32HRPSegwit = "aus"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{23}
	TestNetParams.ScriptHashAddrID = []byte{63}
	TestNetParams.Bech32HRPSegwit = "taus"
}

// AuscashParser handle
type AuscashParser struct {
	*btc.BitcoinParser
}

// NewAuscashParser returns new AuscashParser instance
func NewAuscashParser(params *chaincfg.Params, c *btc.Configuration) *AuscashParser {
	return &AuscashParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Auscash network,
// and the test Auscash network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to litecoin parameters
	// litecoin has dual standard of addresses and we want to be able to
	// parse both standards
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
