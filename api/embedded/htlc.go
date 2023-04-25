package embedded

import (
	"math/big"

	"github.com/ignition-pillar/go-zdk/client"
	"github.com/ignition-pillar/go-zdk/utils/template"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
)

type HtlcApi struct {
	c client.Client
}

func NewHtlcApi(c client.Client) HtlcApi {
	return HtlcApi{c}
}

func (s HtlcApi) GetById(id types.Hash) (*definition.HtlcInfo, error) {
	var result definition.HtlcInfo
	err := s.c.Call(&result, "embedded.htlc.getById", id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s HtlcApi) GetProxyUnlockStatus(address types.Address) (*bool, error) {
	var result bool
	err := s.c.Call(&result, "embedded.htlc.getProxyUnlockStatus", address)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Contract methods

func (s HtlcApi) Create(
	zts types.ZenonTokenStandard,
	amount *big.Int,
	hashLocked types.Address,
	expirationTime int64,
	hashType uint8,
	keyMaxSize uint8,
	hashLock []byte) (*nom.AccountBlock, error) {
	data, err := definition.ABIHtlc.PackMethod(
		definition.CreateHtlcMethodName,
		hashLocked,
		expirationTime,
		hashType,
		keyMaxSize,
		hashLock,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.HtlcContract,
		zts,
		amount,
		data,
	), nil
}

func (s HtlcApi) Reclaim(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABIHtlc.PackMethod(
		definition.ReclaimHtlcMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.HtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (s HtlcApi) Unlock(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABIHtlc.PackMethod(
		definition.UnlockHtlcMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.HtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (s HtlcApi) DenyProxyUnlock() (*nom.AccountBlock, error) {
	data, err := definition.ABIHtlc.PackMethod(
		definition.DenyHtlcProxyUnlockMethodName,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.HtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (s HtlcApi) AllowProxyUnlock() (*nom.AccountBlock, error) {
	data, err := definition.ABIHtlc.PackMethod(
		definition.AllowHtlcProxyUnlockMethodName,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.HtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
