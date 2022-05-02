package client

import "context"

type Client interface {
	Call(result interface{}, method string, args ...interface{}) error
	Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (Subscription, error)
	ProtocolVersion() uint64
	ChainIdentifier() uint64
}

type Subscription interface {
	Err() <-chan error
	Unsubscribe()
}
