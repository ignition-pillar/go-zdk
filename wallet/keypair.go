package wallet

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/wallet"
)

type keyPair struct {
	k *wallet.KeyPair
}

func NewSigner(k *wallet.KeyPair) Signer {
	return &keyPair{k}
}

func (kp *keyPair) Address() types.Address {
	return kp.k.Address
}

func (kp *keyPair) PublicKey() []byte {
	return kp.k.Public
}

func (kp *keyPair) Sign(message []byte) []byte {
	return kp.k.Sign(message)
}

func (kp *keyPair) Signer(data []byte) (signedData []byte, address *types.Address, pubkey []byte, err error) {
	return kp.Sign(data), &kp.k.Address, kp.k.Public, nil
}
