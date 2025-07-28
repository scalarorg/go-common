package chain

import (
	"encoding/binary"
	"fmt"

	"github.com/scalarorg/go-common/types"
)

const ChainInfoBytesSize = 8

const SEPARATOR = "|"

type ChainInfo struct {
	ChainType types.ChainType `json:"chain_type"`
	ChainID   uint64          `json:"chain_id"`
}

func NewChainInfoFromBytes(bytes []byte) *ChainInfo {
	if len(bytes) != ChainInfoBytesSize {
		return nil
	}

	if !types.ValidateChainType(types.ChainType(bytes[0])) {
		return nil
	}
	chainType := types.ChainType(bytes[0])
	chainBytes := make([]byte, len(bytes))
	copy(chainBytes, bytes)
	chainBytes[0] = 0
	chainID := binary.BigEndian.Uint64(chainBytes)

	return &ChainInfo{
		ChainType: chainType,
		ChainID:   chainID,
	}
}

func (dc *ChainInfo) ToBytes() ChainInfoBytes {
	return ChainInfoBytes(dc.Bytes())
}

func (dc *ChainInfo) Bytes() []byte {
	chainIDBytes := make([]byte, ChainInfoBytesSize)
	binary.BigEndian.PutUint64(chainIDBytes, dc.ChainID)

	chainTypeBytes := byte(dc.ChainType)

	bytes := make([]byte, ChainInfoBytesSize)
	bytes[0] = chainTypeBytes
	copy(bytes[1:], chainIDBytes[1:])

	return bytes
}

func (c *ChainInfo) FromString(s string) error {
	chainInfoBytes := ChainInfoBytes{}
	err := chainInfoBytes.FromString(s)
	if err != nil {
		return err
	}

	c.ChainType = chainInfoBytes.ChainType()
	c.ChainID = chainInfoBytes.ChainID()
	return nil
}

func (ChainInfo) Size() int {
	return ChainInfoBytesSize
}

func (c ChainInfo) MarshalTo(data []byte) (int, error) {
	copy(data, c.Bytes())
	return c.Size(), nil
}

func (c *ChainInfo) Unmarshal(data []byte) error {
	if len(data) != c.Size() {
		return fmt.Errorf("invalid data length")
	}
	copy(c.Bytes(), data)
	return nil
}
