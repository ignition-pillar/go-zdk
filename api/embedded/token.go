package embedded

import (
	"math/big"

	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/utils/template"
)

type TokenApi struct {
	client client.Client
}

func NewTokenApi(client client.Client) TokenApi {
	return TokenApi{client}
}

// RPC
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

// Contract methods
func (t *TokenApi) IssueToken(
	tokenName string,
	tokenSymbol string,
	tokenDomain string,
	totalSupply *big.Int,
	maxSupply *big.Int,
	decimals uint8,
	mintable bool,
	burnable bool,
	utility bool) (*nom.AccountBlock, error) {
	data, err := definition.ABIToken.PackMethod(
		definition.IssueMethodName,
		tokenName,
		tokenSymbol,
		tokenDomain,
		totalSupply,
		maxSupply,
		decimals,
		mintable,
		burnable,
		utility,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		t.client.ProtocolVersion(),
		t.client.ChainIdentifier(),
		types.TokenContract,
		types.ZnnTokenStandard,
		constants.TokenIssueAmount,
		data,
	), nil
}

func (t *TokenApi) MintToken(zts types.ZenonTokenStandard, amount *big.Int, receiveAddress types.Address) (*nom.AccountBlock, error) {
	data, err := definition.ABIToken.PackMethod(
		definition.MintMethodName,
		zts,
		amount,
		receiveAddress,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		t.client.ProtocolVersion(),
		t.client.ChainIdentifier(),
		types.TokenContract,
		types.ZnnTokenStandard,
		big.NewInt(0),
		data,
	), nil
}

func (t *TokenApi) BurnToken(zts types.ZenonTokenStandard, amount *big.Int) (*nom.AccountBlock, error) {
	data, err := definition.ABIToken.PackMethod(definition.BurnMethodName)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		t.client.ProtocolVersion(),
		t.client.ChainIdentifier(),
		types.TokenContract,
		zts,
		amount,
		data,
	), nil
}

func (t *TokenApi) UpdateToken(zts types.ZenonTokenStandard, owner types.Address, mintable bool, burnable bool) (*nom.AccountBlock, error) {
	data, err := definition.ABIToken.PackMethod(
		definition.UpdateTokenMethodName,
		zts,
		owner,
		mintable,
		burnable,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		t.client.ProtocolVersion(),
		t.client.ChainIdentifier(),
		types.TokenContract,
		types.ZnnTokenStandard,
		big.NewInt(0),
		data,
	), nil
}
