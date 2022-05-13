package api

import (
	"context"

	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api/subscribe"
	"github.com/zenon-wiki/go-zdk/client"
)

type SubscribeApi struct {
	c client.Client
}

func NewSubscribeApi(c client.Client) SubscribeApi {
	return SubscribeApi{c}
}

func (s SubscribeApi) ToMomentums(ch chan []subscribe.Momentum) (client.Subscription, error) {
	ctx := context.Background()
	sub, err := s.c.Subscribe(ctx, "ledger", ch, "momentums")
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (s SubscribeApi) ToAllAccountBlocks(ch chan []subscribe.AccountBlock) (client.Subscription, error) {
	ctx := context.Background()
	sub, err := s.c.Subscribe(ctx, "ledger", ch, "allAccountBlocks")
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (s SubscribeApi) ToAccountBlocksByAddress(ch chan []subscribe.AccountBlock, address types.Address) (client.Subscription, error) {
	ctx := context.Background()
	sub, err := s.c.Subscribe(ctx, "ledger", ch, "accountBlocksByAddress", address.String())
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (s SubscribeApi) ToUnreceivedAccountBlocksByAddress(ch chan []subscribe.AccountBlock, address types.Address) (client.Subscription, error) {
	ctx := context.Background()
	sub, err := s.c.Subscribe(ctx, "ledger", ch, "unreceivedAccountBlocksByAddress", address.String())
	if err != nil {
		return nil, err
	}
	return sub, nil
}
