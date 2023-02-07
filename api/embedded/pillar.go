package embedded

import (
	"math/big"

	"github.com/ignition-pillar/go-zdk/client"
	"github.com/ignition-pillar/go-zdk/utils/template"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
)

type PillarApi struct {
	c client.Client
}

func NewPillarApi(c client.Client) PillarApi {
	return PillarApi{c}
}

// Common RPC
func (p PillarApi) GetDepositedQsr(address types.Address) (*big.Int, error) {
	var result big.Int
	err := p.c.Call(&result, "embedded.pillar.getDepositedQsr", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) GetUncollectedReward(address types.Address) (*definition.RewardDeposit, error) {
	var result definition.RewardDeposit
	err := p.c.Call(&result, "embedded.pillar.getUncollectedReward", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) GetFrontierRewardByPage(address types.Address, pageIndex, pageSize uint32) (*embedded.RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.RewardHistoryList
	err := p.c.Call(&result, "embedded.pillar.getFrontierRewardByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RPC
func (p PillarApi) GetQsrRegistrationCost() (*big.Int, error) {
	var result big.Int
	err := p.c.Call(&result, "embedded.pillar.getQsrRegistrationCost")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) GetAll(pageIndex, pageSize uint32) (*embedded.PillarInfoList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.PillarInfoList
	err := p.c.Call(&result, "embedded.pillar.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) GetByOwner(address types.Address) ([]embedded.PillarInfo, error) {
	var result []embedded.PillarInfo
	err := p.c.Call(&result, "embedded.pillar.getByOwner", address.String())
	return result, err
}

func (p PillarApi) GetByName(name string) (*embedded.PillarInfo, error) {
	var result embedded.PillarInfo
	err := p.c.Call(&result, "embedded.pillar.getByName", name)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) CheckNameAvailability(name string) (bool, error) {
	var result bool
	err := p.c.Call(&result, "embedded.pillar.checkNameAvailability", name)
	return result, err
}

func (p PillarApi) GetDelegatedPillar(address types.Address) (*embedded.GetDelegatedPillarResponse, error) {
	var result embedded.GetDelegatedPillarResponse
	err := p.c.Call(&result, "embedded.pillar.getDelegatedPillar", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) GetPillarEpochHistory(name string, pageIndex, pageSize uint32) (*embedded.PillarEpochHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.PillarEpochHistoryList
	err := p.c.Call(&result, "embedded.pillar.getPillarEpochHistory", name, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PillarApi) GetPillarsHistoryByEpoch(epoch uint64, pageIndex, pageSize uint32) (*embedded.PillarEpochHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.PillarEpochHistoryList
	err := p.c.Call(&result, "embedded.pillar.getPillarsHistoryByEpoch", epoch, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Contract methods
func (p PillarApi) Register(name string, producerAddress types.Address, rewardAddress types.Address, giveBlockRewardPercentage uint8, giveDelegateRewardPercentage uint8) (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(
		definition.RegisterMethodName,
		name,
		producerAddress,
		rewardAddress,
		giveBlockRewardPercentage,
		giveDelegateRewardPercentage,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		constants.PillarStakeAmount,
		data,
	), nil
}

func (p PillarApi) RegisterLegacy(name string, producerAddress types.Address, rewardAddress types.Address, publicKey string, signature string, giveBlockRewardPercentage uint8, giveDelegateRewardPercentage uint8) (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(
		definition.LegacyRegisterMethodName,
		name,
		producerAddress,
		rewardAddress,
		giveBlockRewardPercentage,
		giveDelegateRewardPercentage,
		publicKey,
		signature,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		constants.PillarStakeAmount,
		data,
	), nil
}

func (p PillarApi) UpdatePillar(name string, producerAddress types.Address, rewardAddress types.Address, giveBlockRewardPercentage uint8, giveDelegateRewardPercentage uint8) (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(
		definition.UpdatePillarMethodName,
		name,
		producerAddress,
		rewardAddress,
		giveBlockRewardPercentage,
		giveDelegateRewardPercentage,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (p PillarApi) Revoke(name string) (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(
		definition.RevokeMethodName,
		name,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (p PillarApi) Delegate(name string) (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(
		definition.DelegateMethodName,
		name,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (p PillarApi) Undelegate() (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(definition.UndelegateMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

// Common contract methods
func (p PillarApi) CollectReward() (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(definition.CollectRewardMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (p PillarApi) DepositQsr(amount *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(definition.DepositQsrMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.QsrTokenStandard,
		amount,
		data,
	), nil
}

func (p PillarApi) WithdrawQsr() (*nom.AccountBlock, error) {
	data, err := definition.ABIPillars.PackMethod(definition.WithdrawQsrMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PillarContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
