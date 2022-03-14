package embedded

import (
	"math/big"

	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
)

type SentinelApi struct {
	client client.IClient
}

func (s *SentinelApi) SetClient(client client.IClient) {
	s.client = client
}

// RPC
func (s *SentinelApi) GetAllActive(pageIndex, pageSize uint32) (*embedded.SentinelInfoList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.SentinelInfoList
	err := s.client.Call(&result, "embedded.sentinel.getAllActive", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SentinelApi) GetByOwner(address types.Address) (*embedded.SentinelInfo, error) {
	var result embedded.SentinelInfo
	err := s.client.Call(&result, "embedded.sentinel.getByOwner", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Common RPC
func (s *SentinelApi) GetDepositedQsr(address types.Address) (*big.Int, error) {
	var result big.Int
	err := s.client.Call(&result, "embedded.sentinel.getDepositedQsr", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SentinelApi) GetUncollectedReward(address types.Address) (*definition.RewardDeposit, error) {
	var result definition.RewardDeposit
	err := s.client.Call(&result, "embedded.sentinel.getUncollectedReward", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SentinelApi) GetFrontierRewardByPage(address types.Address, pageIndex, pageSize uint32) (*embedded.RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.RewardHistoryList
	err := s.client.Call(&result, "embedded.sentinel.getFrontierRewardByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
