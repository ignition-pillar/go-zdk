package embedded

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-wiki/go-zdk/client"
)

type PlasmaApi struct {
	client client.IClient
}

func (p *PlasmaApi) SetClient(client client.IClient) {
	p.client = client
}

func (p *PlasmaApi) Get(address types.Address) (*embedded.PlasmaInfo, error) {
	var result embedded.PlasmaInfo
	err := p.client.Call(&result, "embedded.plasma.get", address.String())
	return &result, err
}

func (p *PlasmaApi) GetEntriesByAddress(address types.Address, pageIndex, pageSize uint32) (*embedded.FusionEntryList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.FusionEntryList
	err := p.client.Call(&result, "embedded.plasma.getEntriesByAddress", address.String(), pageIndex, pageSize)
	return &result, err
}

func (p *PlasmaApi) GetRequiredPoWForAccountBlock(param *embedded.GetRequiredParam) (*embedded.GetRequiredResult, error) {
	var result embedded.GetRequiredResult
	err := p.client.Call(&result, "embedded.plasma.getRequiredPoWForAccountBlock", param)
	return &result, err
}
