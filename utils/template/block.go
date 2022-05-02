package template

import (
	"math/big"

	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
)

func Receive(version uint64, chainId uint64, from types.Hash) *nom.AccountBlock {
	block := nom.AccountBlock{}
	block.Version = version
	block.ChainIdentifier = chainId
	block.BlockType = nom.BlockTypeUserReceive
	block.FromBlockHash = from
	return &block
}

func Send(version uint64, chainId uint64, to types.Address, zts types.ZenonTokenStandard, amount *big.Int, data []byte) *nom.AccountBlock {
	block := nom.AccountBlock{}
	block.Version = version
	block.ChainIdentifier = chainId
	block.BlockType = nom.BlockTypeUserSend
	block.ToAddress = to
	block.TokenStandard = zts
	block.Amount = amount
	block.Data = data
	return &block
}

var CallContract = Send
