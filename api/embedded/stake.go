package embedded

import (
	"math/big"

	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/utils/template"
)

type StakeApi struct {
	c client.Client
}

func NewStakeApi(c client.Client) StakeApi {
	return StakeApi{c}
}

// RPC
func (s StakeApi) GetEntriesByAddress(address types.Address, pageIndex, pageSize uint32) (*embedded.StakeList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.StakeList
	err := s.c.Call(&result, "embedded.stake.getEntriesByAddress", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Common RPC
func (s StakeApi) GetUncollectedReward(address types.Address) (*definition.RewardDeposit, error) {
	var result definition.RewardDeposit
	err := s.c.Call(&result, "embedded.stake.getUncollectedReward", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s StakeApi) GetFrontierRewardByPage(address types.Address, pageIndex, pageSize uint32) (*embedded.RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.RewardHistoryList
	err := s.c.Call(&result, "embedded.stake.getFrontierRewardByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Contract methods
func (s StakeApi) Stake(durationInSec int64, amount *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIStake.PackMethod(
		definition.StakeMethodName,
		durationInSec,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.StakeContract,
		types.ZnnTokenStandard,
		amount,
		data,
	), nil
}

func (s StakeApi) Cancel(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABIStake.PackMethod(
		definition.CancelStakeMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.StakeContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

// Common contract methods
func (s StakeApi) CollectReward() (*nom.AccountBlock, error) {
	data, err := definition.ABIStake.PackMethod(definition.CollectRewardMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.StakeContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
