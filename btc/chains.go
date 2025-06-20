package btc

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/scalarorg/go-common/types"
)

var chaincfgTestnet4ParamsName = "testnet4"

// Chain represents a Bitcoin network configuration
type chainType struct {
	*types.BaseChain
	Params     *chaincfg.Params
	ParamsName string
}

var (
	displayedNameByID map[uint64]string
	paramsByName      map[string]*chaincfg.Params
)

func init() {
	displayedNameByID = make(map[uint64]string)
	paramsByName = make(map[string]*chaincfg.Params)

	for _, chain := range chains {
		displayedNameByID[chain.ID] = chain.DisplayedName
		paramsByName[chain.ParamsName] = chain.Params
	}
}

// BTCChainRecords provides access to Bitcoin chain configuration and parameters
type Records struct{}

var _ types.ChainRecordsType = &Records{}

var records = &Records{}

func BtcChainsRecords() *Records {
	return records
}

// GetChainParamsByName returns the chain parameters for a given chain name.
// Eg: "mainnet", "testnet3", "testnet4", "regression", "signet"
func (b *Records) GetChainParamsByName(paramsName string) *chaincfg.Params {
	params, ok := paramsByName[paramsName]
	if !ok {
		return nil
	}
	return params
}

// GetDisplayedName returns the displayed name for a given chain ID.
// Returns empty string if the chain ID is not found.
func (b *Records) GetDisplayedName(chainID uint64) string {
	name, ok := displayedNameByID[chainID]
	if !ok {
		return ""
	}
	return name
}

func (b *Records) GetChainParamsByID(chainID uint64) *chaincfg.Params {
	for _, chain := range chains {
		if chain.ID == chainID {
			return chain.Params
		}
	}
	return nil
}
