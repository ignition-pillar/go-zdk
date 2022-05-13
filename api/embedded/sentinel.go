package embedded

import (
	"math/big"

	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/utils/template"
)

type SentinelApi struct {
	c client.Client
}

func NewSentinelApi(c client.Client) SentinelApi {
	return SentinelApi{c}
}

// RPC
func (s SentinelApi) GetAllActive(pageIndex, pageSize uint32) (*embedded.SentinelInfoList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.SentinelInfoList
	err := s.c.Call(&result, "embedded.sentinel.getAllActive", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s SentinelApi) GetByOwner(address types.Address) (*embedded.SentinelInfo, error) {
	var result embedded.SentinelInfo
	err := s.c.Call(&result, "embedded.sentinel.getByOwner", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Common RPC
func (s SentinelApi) GetDepositedQsr(address types.Address) (*big.Int, error) {
	var result big.Int
	err := s.c.Call(&result, "embedded.sentinel.getDepositedQsr", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s SentinelApi) GetUncollectedReward(address types.Address) (*definition.RewardDeposit, error) {
	var result definition.RewardDeposit
	err := s.c.Call(&result, "embedded.sentinel.getUncollectedReward", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s SentinelApi) GetFrontierRewardByPage(address types.Address, pageIndex, pageSize uint32) (*embedded.RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.RewardHistoryList
	err := s.c.Call(&result, "embedded.sentinel.getFrontierRewardByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Contract methods
func (s SentinelApi) Register() (*nom.AccountBlock, error) {
	data, err := definition.ABISentinel.PackMethod(definition.RegisterSentinelMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.SentinelContract,
		types.ZnnTokenStandard,
		constants.SentinelZnnRegisterAmount,
		data,
	), nil
}

func (s SentinelApi) Revoke() (*nom.AccountBlock, error) {
	data, err := definition.ABISentinel.PackMethod(definition.RevokeSentinelMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.SentinelContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

// Common contract methods
func (s SentinelApi) CollectReward() (*nom.AccountBlock, error) {
	data, err := definition.ABISentinel.PackMethod(definition.CollectRewardMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.SentinelContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (s SentinelApi) DepositQsr(amount *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABISentinel.PackMethod(definition.DepositQsrMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.SentinelContract,
		types.QsrTokenStandard,
		amount,
		data,
	), nil
}

func (s SentinelApi) WithdrawQsr() (*nom.AccountBlock, error) {
	data, err := definition.ABISentinel.PackMethod(definition.WithdrawQsrMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.SentinelContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
