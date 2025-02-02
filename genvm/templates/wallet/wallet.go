package wallet

import (
	"fmt"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

// New returns Wallet instance with SpawnArguments.
func New(args *SpawnArguments) *Wallet {
	return &Wallet{PublicKey: args.PublicKey}
}

//go:generate scalegen

// Wallet is a single-key wallet.
type Wallet struct {
	PublicKey core.PublicKey
}

// MaxSpend returns amount specified in the SpendArguments for Spend method.
func (s *Wallet) MaxSpend(method uint8, args any) (uint64, error) {
	switch method {
	case methodSpawn:
		return 0, nil
	case methodSpend:
		return args.(*SpendArguments).Amount, nil
	default:
		return 0, fmt.Errorf("%w: unknown method %d", core.ErrMalformed, method)
	}
}

// Verify that transaction is signed by the owner of the PublicKey using ed25519.
func (s *Wallet) Verify(ctx *core.Context, raw []byte, dec *scale.Decoder) bool {
	sig := core.Signature{}
	n, err := sig.DecodeScale(dec)
	if err != nil {
		return false
	}
	hash := core.Hash(raw[:len(raw)-n])
	return ed25519.Verify(ed25519.PublicKey(s.PublicKey[:]), hash[:], sig[:])
}

// Spend transfers an amount to the address specified in SpendArguments.
func (s *Wallet) Spend(ctx *core.Context, args *SpendArguments) error {
	return ctx.Transfer(args.Destination, args.Amount)
}
