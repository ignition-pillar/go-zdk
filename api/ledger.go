package api

import (
	"github.com/ignition-pillar/go-zdk/client"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api"
)

type LedgerApi struct {
	c client.Client
}

func NewLedgerApi(c client.Client) LedgerApi {
	return LedgerApi{c}
}

const (
	unreceivedMaxPageIndex = 10
	unreceivedMaxPageSize  = 50
)

func (l LedgerApi) PublishRawTransaction(block *nom.AccountBlock) error {
	err := l.c.Call(nil, "ledger.publishRawTransaction", block)
	return err
}

func (l LedgerApi) GetUnconfirmedBlocksByAddress(address types.Address, pageIndex, pageSize uint32) (*api.AccountBlockList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result api.AccountBlockList
	err := l.c.Call(&result, "ledger.getUnconfirmedBlocksByAddress", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetUnreceivedBlocksByAddress(address types.Address, pageIndex, pageSize uint32) (*api.AccountBlockList, error) {
	if pageIndex > unreceivedMaxPageIndex {
		pageIndex = unreceivedMaxPageIndex
	}
	if pageSize > unreceivedMaxPageSize {
		pageSize = unreceivedMaxPageSize
	}
	var result api.AccountBlockList
	err := l.c.Call(&result, "ledger.getUnreceivedBlocksByAddress", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Blocks
func (l LedgerApi) GetFrontierAccountBlock(address types.Address) (*api.AccountBlock, error) {
	var result api.AccountBlock
	err := l.c.Call(&result, "ledger.getFrontierAccountBlock", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetAccountBlockByHash(hash types.Hash) (*api.AccountBlock, error) {
	var result api.AccountBlock
	err := l.c.Call(&result, "ledger.getAccountBlockByHash", hash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetAccountBlocksByHeight(address types.Address, height, count uint64) (*api.AccountBlockList, error) {
	if height == 0 {
		height = 1
	}
	if count > api.RpcMaxPageSize {
		count = api.RpcMaxPageSize
	}
	var result api.AccountBlockList
	err := l.c.Call(&result, "ledger.getAccountBlocksByHeight", address.String(), height, count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// pageIndex = 0 returns the most recent account blocks sorted descending by height
func (l LedgerApi) GetAccountBlocksByPage(address types.Address, pageIndex, pageSize uint32) (*api.AccountBlockList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result api.AccountBlockList
	err := l.c.Call(&result, "ledger.getAccountBlocksByPage", address.String(), pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Momentums
func (l LedgerApi) GetFrontierMomentum() (*api.Momentum, error) {
	var result api.Momentum
	err := l.c.Call(&result, "ledger.getFrontierMomentum")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetMomentumBeforeTime(timestamp uint64) (*api.Momentum, error) {
	var result api.Momentum
	err := l.c.Call(&result, "ledger.getMomentumBeforeTime", timestamp)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetMomentumByHash(hash types.Hash) (*api.Momentum, error) {
	var result api.Momentum
	err := l.c.Call(&result, "ledger.getMomentumByHash", hash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetMomentumsByHeight(height, count uint64) (*api.MomentumList, error) {
	if height == 0 {
		height = 1
	}
	if count > api.RpcMaxPageSize {
		count = api.RpcMaxPageSize
	}
	var result api.MomentumList
	err := l.c.Call(&result, "ledger.getMomentumsByHeight", height, count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// pageIndex = 0 returns the most recent momentums sorted descending by height
func (l LedgerApi) GetMomentumsByPage(pageIndex, pageSize uint32) (*api.MomentumList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result api.MomentumList
	err := l.c.Call(&result, "ledger.getMomentumsByPage", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (l LedgerApi) GetDetailedMomentumsByHeight(height, count uint64) (*api.DetailedMomentumList, error) {
	if height == 0 {
		height = 1
	}
	if count > api.RpcMaxPageSize {
		count = api.RpcMaxPageSize
	}
	var result api.DetailedMomentumList
	err := l.c.Call(&result, "ledger.getDetailedMomentumsByHeight", height, count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Account info
func (l LedgerApi) GetAccountInfoByAddress(address types.Address) (*api.AccountInfo, error) {
	var result api.AccountInfo
	err := l.c.Call(&result, "ledger.getAccountInfoByAddress", address.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}
