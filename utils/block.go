package utils

import (
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/ignition-pillar/go-zdk/wallet"
	"github.com/ignition-pillar/go-zdk/zdk"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/pow"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
)

func autofillTransactionParameters(z *zdk.Zdk, tx *nom.AccountBlock) error {
	frontierAccountBlock, err := z.Ledger.GetFrontierAccountBlock(tx.Address)
	if err != nil {
		return err
	}
	var height uint64 = 1
	previousHash := types.ZeroHash
	if frontierAccountBlock != nil {
		height = frontierAccountBlock.Height + 1
		previousHash = frontierAccountBlock.Hash
	}

	tx.Height = height
	tx.PreviousHash = previousHash

	frontierMomentum, err := z.Ledger.GetFrontierMomentum()
	if err != nil {
		return err
	}
	momentumAcknowledged := types.HashHeight{Hash: frontierMomentum.Hash, Height: frontierMomentum.Height}
	tx.MomentumAcknowledged = momentumAcknowledged

	return nil
}

func checkAndSetFields(z *zdk.Zdk, tx *nom.AccountBlock, signer wallet.Signer) error {
	tx.Address = signer.Address()
	tx.PublicKey = signer.PublicKey()

	err := autofillTransactionParameters(z, tx)
	if err != nil {
		return err
	}

	// Use IsReceiveBlock() instead?
	if tx.IsSendBlock() {
	} else {
		if tx.FromBlockHash == types.ZeroHash {
			return errors.New("fromblockhash cannot be zero")
		}

		sendBlock, err := z.Ledger.GetAccountBlockByHash(tx.FromBlockHash)
		if err != nil {
			return err
		}
		if sendBlock == nil {
			return errors.New("sendblock does not exist")
		}

		if sendBlock.ToAddress != tx.Address {
			return errors.New("signer address does not match sendblock")
		}

		if len(tx.Data) != 0 {
			return errors.New("receive blocks must have empty data")
		}

		// is there a better way to check? just cast as string?
		if tx.Difficulty > 0 && hex.EncodeToString(tx.Nonce.Data[:]) == "" {
			return errors.New("invalid pow")
		}
	}

	return nil
}

// TODO support powCallback
func setDifficulty(z *zdk.Zdk, tx *nom.AccountBlock, waitForRequiredPlasma bool) error {
	param := &embedded.GetRequiredParam{
		SelfAddr:  tx.Address,
		BlockType: tx.BlockType,
		ToAddr:    &tx.ToAddress,
		Data:      tx.Data,
	}
	resp, err := z.Embedded.Plasma.GetRequiredPoWForAccountBlock(param)
	if err != nil {
		return err
	}

	if resp.RequiredDifficulty != 0 {
		tx.FusedPlasma = resp.AvailablePlasma
		tx.Difficulty = resp.RequiredDifficulty
		dataHash := pow.GetAccountBlockHash(tx)
		var d big.Int
		d.SetUint64(resp.RequiredDifficulty)
		tx.Nonce = nom.DeSerializeNonce(pow.GetPoWNonce(&d, dataHash))
	} else {
		tx.FusedPlasma = resp.BasePlasma
		tx.Difficulty = 0
		tx.Nonce = nom.Nonce{}
	}
	return nil
}

func setHashAndSignature(tx *nom.AccountBlock, signer wallet.Signer) {
	tx.Hash = tx.ComputeHash()
	tx.Signature = signer.Sign(tx.Hash.Bytes())
}

// TODO support powCallback
func Send(z *zdk.Zdk, tx *nom.AccountBlock, signer wallet.Signer, waitForRequiredPlasma bool) (*nom.AccountBlock, error) {
	err := checkAndSetFields(z, tx, signer)
	if err != nil {
		return nil, err
	}
	err = setDifficulty(z, tx, waitForRequiredPlasma)
	if err != nil {
		return nil, err
	}

	setHashAndSignature(tx, signer)

	err = z.Ledger.PublishRawTransaction(tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func RequiresPow(z *zdk.Zdk, tx *nom.AccountBlock, signer wallet.Signer) (*bool, error) {
	tx.Address = signer.Address()
	powParam := &embedded.GetRequiredParam{
		SelfAddr:  tx.Address,
		BlockType: tx.BlockType,
		ToAddr:    &tx.ToAddress,
		Data:      tx.Data,
	}
	var ans bool = true
	resp, err := z.Embedded.Plasma.GetRequiredPoWForAccountBlock(powParam)
	if err != nil {
		return nil, err
	}
	if resp.RequiredDifficulty == 0 {
		ans = false
	}
	return &ans, nil
}
