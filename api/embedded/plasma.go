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
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
)

type PlasmaApi struct {
	c client.Client
}

func NewPlasmaApi(c client.Client) PlasmaApi {
	return PlasmaApi{c}
}

func (p PlasmaApi) Get(address types.Address) (*embedded.PlasmaInfo, error) {
	var result embedded.PlasmaInfo
	err := p.c.Call(&result, "embedded.plasma.get", address.String())
	return &result, err
}

func (p PlasmaApi) GetEntriesByAddress(address types.Address, pageIndex, pageSize uint32) (*embedded.FusionEntryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.FusionEntryList
	err := p.c.Call(&result, "embedded.plasma.getEntriesByAddress", address.String(), pageIndex, pageSize)
	return &result, err
}

func (p PlasmaApi) GetRequiredPoWForAccountBlock(param *embedded.GetRequiredParam) (*embedded.GetRequiredResult, error) {
	var result embedded.GetRequiredResult
	err := p.c.Call(&result, "embedded.plasma.getRequiredPoWForAccountBlock", param)
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
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
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
		p.c.ProtocolVersion(),
		p.c.ChainIdentifier(),
		types.PlasmaContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
