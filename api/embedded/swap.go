package embedded

import (
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/utils/template"
)

type SwapApi struct {
	c client.Client
}

func NewSwapApi(c client.Client) SwapApi {
	return SwapApi{c}
}

// RPC
func (s SwapApi) GetAssetsByKeyIdHash(keyIdHash types.Hash) (*embedded.SwapAssetEntry, error) {
	var result embedded.SwapAssetEntry
	err := s.c.Call(&result, "embedded.swap.getAssetsByKeyIdHash", keyIdHash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s SwapApi) GetAssets() (map[types.Hash]embedded.SwapAssetEntrySimple, error) {
	var result map[types.Hash]embedded.SwapAssetEntrySimple
	err := s.c.Call(&result, "embedded.swap.getAssets")
	return result, err
}

func (s SwapApi) GetLegacyPillars() ([]embedded.SwapLegacyPillarEntry, error) {
	var result []embedded.SwapLegacyPillarEntry
	err := s.c.Call(&result, "embedded.swap.getLegacyPillars")
	return result, err
}

// Contract methods
func (s SwapApi) RetrieveAssets(pubKey string, signature string) (*nom.AccountBlock, error) {
	data, err := definition.ABISwap.PackMethod(
		definition.RetrieveAssetsMethodName,
		pubKey,
		signature,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.SwapContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
