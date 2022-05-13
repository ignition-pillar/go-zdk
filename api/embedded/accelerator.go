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

type AcceleratorApi struct {
	c client.Client
}

func NewAcceleratorApi(c client.Client) AcceleratorApi {
	return AcceleratorApi{c}
}

// RPC
func (a AcceleratorApi) GetAll(pageIndex, pageSize uint32) (*embedded.ProjectList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.ProjectList
	err := a.c.Call(&result, "embedded.accelerator.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a AcceleratorApi) GetProjectById(id types.Hash) (*embedded.Project, error) {
	var result embedded.Project
	err := a.c.Call(&result, "embedded.accelerator.getProjectById", id.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a AcceleratorApi) GetPhaseById(id types.Hash) (*embedded.Phase, error) {
	var result embedded.Phase
	err := a.c.Call(&result, "embedded.accelerator.getPhaseById", id.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a AcceleratorApi) GetPillarVotes(name string, hashes []types.Hash) ([]definition.PillarVote, error) {
	var result []definition.PillarVote
	err := a.c.Call(&result, "embedded.accelerator.getPillarVotes", name, hashes)
	return result, err
}

func (a AcceleratorApi) GetVoteBreakdown(id types.Hash) (*definition.VoteBreakdown, error) {
	var result definition.VoteBreakdown
	err := a.c.Call(&result, "embedded.accelerator.getVoteBreakdown", id.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ContractMethods
func (a AcceleratorApi) CreateProject(name string, description string, url string, znnNeeded *big.Int, qsrNeeded *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIAccelerator.PackMethod(
		definition.CreateProjectMethodName,
		name,
		description,
		url,
		znnNeeded,
		qsrNeeded,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		a.c.ProtocolVersion(),
		a.c.ChainIdentifier(),
		types.AcceleratorContract,
		types.ZnnTokenStandard,
		constants.ProjectCreationAmount,
		data,
	), nil
}

func (a AcceleratorApi) AddPhase(id types.Hash, name string, description string, url string, znnNeeded *big.Int, qsrNeeded *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIAccelerator.PackMethod(
		definition.AddPhaseMethodName,
		id,
		name,
		description,
		url,
		znnNeeded,
		qsrNeeded,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		a.c.ProtocolVersion(),
		a.c.ChainIdentifier(),
		types.AcceleratorContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (a AcceleratorApi) UpdatePhase(id types.Hash, name string, description string, url string, znnNeeded *big.Int, qsrNeeded *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIAccelerator.PackMethod(
		definition.UpdatePhaseMethodName,
		id,
		name,
		description,
		url,
		znnNeeded,
		qsrNeeded,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		a.c.ProtocolVersion(),
		a.c.ChainIdentifier(),
		types.AcceleratorContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (a AcceleratorApi) Donate(amount *big.Int, zts types.ZenonTokenStandard) (*nom.AccountBlock, error) {
	data, err := definition.ABIAccelerator.PackMethod(definition.DonateMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		a.c.ProtocolVersion(),
		a.c.ChainIdentifier(),
		types.AcceleratorContract,
		zts,
		amount,
		data,
	), nil
}

func (a AcceleratorApi) VoteByName(id types.Hash, pillarName string, vote uint8) (*nom.AccountBlock, error) {
	data, err := definition.ABIAccelerator.PackMethod(
		definition.VoteByNameMethodName,
		id,
		pillarName,
		vote,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		a.c.ProtocolVersion(),
		a.c.ChainIdentifier(),
		types.AcceleratorContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (a AcceleratorApi) VoteByProdAddress(id types.Hash, vote uint8) (*nom.AccountBlock, error) {
	data, err := definition.ABIAccelerator.PackMethod(
		definition.VoteByProdAddressMethodName,
		id,
		vote,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		a.c.ProtocolVersion(),
		a.c.ChainIdentifier(),
		types.AcceleratorContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
