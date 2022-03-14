package embedded

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
)

type StakeApi struct {
	client client.IClient
}

func (s *StakeApi) SetClient(client client.IClient) {
	s.client = client
}

// RPC
func (s *StakeApi) GetEntriesByAddress(address types.Address, pageIndex, pageSize uint32) (*embedded.StakeList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.StakeList
	err := s.client.Call(&result, "embedded.stake.getEntriesByAddress", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Common RPC
func (s *StakeApi) GetUncollectedReward(address types.Address) (*definition.RewardDeposit, error) {
	var result definition.RewardDeposit
	err := s.client.Call(&result, "embedded.stake.getUncollectedReward", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *StakeApi) GetFrontierRewardByPage(address types.Address, pageIndex, pageSize uint32) (*embedded.RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.RewardHistoryList
	err := s.client.Call(&result, "embedded.stake.getFrontierRewardByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
