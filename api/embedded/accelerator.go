package embedded

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
)

type AcceleratorApi struct {
	client client.IClient
}

func (a *AcceleratorApi) SetClient(client client.IClient) {
	a.client = client
}

func (a *AcceleratorApi) GetAll(pageIndex, pageSize uint32) (*embedded.ProjectList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.ProjectList
	err := a.client.Call(&result, "embedded.accelerator.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *AcceleratorApi) GetProjectById(id types.Hash) (*embedded.Project, error) {
	var result embedded.Project
	err := a.client.Call(&result, "embedded.accelerator.getProjectById", id.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *AcceleratorApi) GetPhaseById(id types.Hash) (*embedded.Phase, error) {
	var result embedded.Phase
	err := a.client.Call(&result, "embedded.accelerator.getPhaseById", id.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *AcceleratorApi) GetVoteBreakdown(id types.Hash) (*definition.VoteBreakdown, error) {
	var result definition.VoteBreakdown
	err := a.client.Call(&result, "embedded.accelerator.getVoteBreakdown", id.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *AcceleratorApi) GetPillarVotes(name string, hashes []types.Hash) ([]definition.PillarVote, error) {
	var result []definition.PillarVote
	err := a.client.Call(&result, "embedded.accelerator.getPillarVotes", name, hashes)
	return result, err
}
