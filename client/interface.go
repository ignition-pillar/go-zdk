package client

import "context"

type IClient interface {
	Call(result interface{}, method string, args ...interface{}) error
	Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (ISubscription, error)
}

type ISubscription interface {
	Err() <-chan error
	Unsubscribe()
}
