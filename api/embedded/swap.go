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
	client client.Client
}

func NewSwapApi(client client.Client) SwapApi {
	return SwapApi{client}
}

// RPC
func (s SwapApi) GetAssetsByKeyIdHash(keyIdHash types.Hash) (*embedded.SwapAssetEntry, error) {
	var result embedded.SwapAssetEntry
	err := s.client.Call(&result, "embedded.swap.getAssetsByKeyIdHash", keyIdHash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s SwapApi) GetAssets() (map[types.Hash]embedded.SwapAssetEntrySimple, error) {
	var result map[types.Hash]embedded.SwapAssetEntrySimple
	err := s.client.Call(&result, "embedded.swap.getAssets")
	return result, err
}

func (s SwapApi) GetLegacyPillars() ([]embedded.SwapLegacyPillarEntry, error) {
	var result []embedded.SwapLegacyPillarEntry
	err := s.client.Call(&result, "embedded.swap.getLegacyPillars")
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
		s.client.ProtocolVersion(),
		s.client.ChainIdentifier(),
		types.SwapContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
