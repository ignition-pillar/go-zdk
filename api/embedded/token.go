package embedded

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-wiki/go-zdk/client"
)

type TokenApi struct {
	client client.IClient
}

func (t *TokenApi) SetClient(client client.IClient) {
	t.client = client
}

func (t *TokenApi) GetAll(pageIndex, pageSize uint32) (*embedded.TokenList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.TokenList
	err := t.client.Call(&result, "embedded.token.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TokenApi) GetByOwner(address types.Address, pageIndex, pageSize uint32) (*embedded.TokenList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.TokenList
	err := t.client.Call(&result, "embedded.token.getByOwner", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TokenApi) GetByZts(zts types.ZenonTokenStandard) (*api.Token, error) {
	var result api.Token
	err := t.client.Call(&result, "embedded.token.getByZts", zts.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}
