package embedded

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-wiki/go-zdk/client"
)

type SwapApi struct {
	client client.IClient
}

func (s *SwapApi) SetClient(client client.IClient) {
	s.client = client
}

// RPC
func (s *SwapApi) GetAssetsByKeyIdHash(keyIdHash types.Hash) (*embedded.SwapAssetEntry, error) {
	var result embedded.SwapAssetEntry
	err := s.client.Call(&result, "embedded.swap.getAssetsByKeyIdHash", keyIdHash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SwapApi) GetAssets() (map[types.Hash]embedded.SwapAssetEntrySimple, error) {
	var result map[types.Hash]embedded.SwapAssetEntrySimple
	err := s.client.Call(&result, "embedded.swap.getAssets")
	return result, err
}

func (s *SwapApi) GetLegacyPillars() ([]embedded.SwapLegacyPillarEntry, error) {
	var result []embedded.SwapLegacyPillarEntry
	err := s.client.Call(&result, "embedded.swap.getLegacyPillars")
	return result, err
}
