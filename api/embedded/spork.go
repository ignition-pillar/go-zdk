package embedded

import (
	"math/big"

	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/utils/template"
)

type SporkApi struct {
	client client.Client
}

func NewSporkApi(client client.Client) SporkApi {
	return SporkApi{client}
}

func (s SporkApi) GetAll(pageIndex, pageSize uint32) (*embedded.SporkList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.SporkList
	err := s.client.Call(&result, "embedded.spork.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Contract methods
func (s SporkApi) Create(name string, description string) (*nom.AccountBlock, error) {
	data, err := definition.ABISpork.PackMethod(
		definition.SporkCreateMethodName,
		name,
		description,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.client.ProtocolVersion(),
		s.client.ChainIdentifier(),
		types.SporkContract,
		types.ZnnTokenStandard,
		big.NewInt(0),
		data,
	), nil
}

func (s SporkApi) Activate(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABISpork.PackMethod(
		definition.SporkActivateMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.client.ProtocolVersion(),
		s.client.ChainIdentifier(),
		types.SporkContract,
		types.ZnnTokenStandard,
		big.NewInt(0),
		data,
	), nil
}
