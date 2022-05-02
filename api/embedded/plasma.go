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

type PlasmaApi struct {
	client client.Client
}

func NewPlasmaApi(client client.Client) PlasmaApi {
	return PlasmaApi{client}
}

func (p PlasmaApi) Get(address types.Address) (*embedded.PlasmaInfo, error) {
	var result embedded.PlasmaInfo
	err := p.client.Call(&result, "embedded.plasma.get", address.String())
	return &result, err
}

func (p PlasmaApi) GetEntriesByAddress(address types.Address, pageIndex, pageSize uint32) (*embedded.FusionEntryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.FusionEntryList
	err := p.client.Call(&result, "embedded.plasma.getEntriesByAddress", address.String(), pageIndex, pageSize)
	return &result, err
}

func (p PlasmaApi) GetRequiredPoWForAccountBlock(param *embedded.GetRequiredParam) (*embedded.GetRequiredResult, error) {
	var result embedded.GetRequiredResult
	err := p.client.Call(&result, "embedded.plasma.getRequiredPoWForAccountBlock", param)
	return &result, err
}

// Contract methods
func (p PlasmaApi) Fuse(beneficiary types.Address, amount *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIPlasma.PackMethod(
		definition.FuseMethodName,
		beneficiary,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.client.ProtocolVersion(),
		p.client.ChainIdentifier(),
		types.PlasmaContract,
		types.QsrTokenStandard,
		amount,
		data,
	), nil
}

func (p PlasmaApi) Cancel(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABIPlasma.PackMethod(
		definition.CancelFuseMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		p.client.ProtocolVersion(),
		p.client.ChainIdentifier(),
		types.PlasmaContract,
		types.ZnnTokenStandard,
		big.NewInt(0),
		data,
	), nil
}
