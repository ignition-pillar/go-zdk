package wallet

import (
	"github.com/zenon-network/go-zenon/common/types"
)

type Signer interface {
	Address() types.Address
	PublicKey() []byte
	Sign(message []byte) []byte
	Signer(data []byte) (signedData []byte, address *types.Address, pubkey []byte, err error)
}
