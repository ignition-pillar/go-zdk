package embedded

import (
	"math/big"

	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
)

type PillarApi struct {
	client client.IClient
}

func (p *PillarApi) SetClient(client client.IClient) {
	p.client = client
}

// Common RPC
func (p *PillarApi) GetDepositedQsr(address types.Address) (*big.Int, error) {
	var result big.Int
	err := p.client.Call(&result, "embedded.pillar.getDepositedQsr", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) GetUncollectedReward(address types.Address) (*definition.RewardDeposit, error) {
	var result definition.RewardDeposit
	err := p.client.Call(&result, "embedded.pillar.getUncollectedReward", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) GetFrontierRewardByPage(address types.Address, pageIndex, pageSize uint32) (*embedded.RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.RewardHistoryList
	err := p.client.Call(&result, "embedded.pillar.getFrontierRewardByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RPC
func (p *PillarApi) GetQsrRegistrationCost() (*big.Int, error) {
	var result big.Int
	err := p.client.Call(&result, "embedded.pillar.getQsrRegistrationCost")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) GetAll(pageIndex, pageSize uint32) (*embedded.PillarInfoList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.PillarInfoList
	err := p.client.Call(&result, "embedded.pillar.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) GetByOwner(address types.Address) ([]embedded.PillarInfo, error) {
	var result []embedded.PillarInfo
	err := p.client.Call(&result, "embedded.pillar.getByOwner", address.String())
	return result, err
}

func (p *PillarApi) GetByName(name string) (*embedded.PillarInfo, error) {
	var result embedded.PillarInfo
	err := p.client.Call(&result, "embedded.pillar.getByName", name)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) CheckNameAvailability(name string) (bool, error) {
	var result bool
	err := p.client.Call(&result, "embedded.pillar.checkNameAvailability", name)
	return result, err
}

func (p *PillarApi) GetDelegatedPillar(address types.Address) (*embedded.GetDelegatedPillarResponse, error) {
	var result embedded.GetDelegatedPillarResponse
	err := p.client.Call(&result, "embedded.pillar.getDelegatedPillar", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) GetPillarEpochHistory(name string, pageIndex, pageSize uint32) (*embedded.PillarEpochHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.PillarEpochHistoryList
	err := p.client.Call(&result, "embedded.pillar.getPillarEpochHistory", name, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PillarApi) GetPillarsHistoryByEpoch(epoch uint64, pageIndex, pageSize uint32) (*embedded.PillarEpochHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.PillarEpochHistoryList
	err := p.client.Call(&result, "embedded.pillar.getPillarsHistoryByEpoch", epoch, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
