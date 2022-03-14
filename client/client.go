package client

import (
	"context"

	rpc "github.com/zenon-network/go-zenon/rpc/server"
)

type Client struct {
	ec *rpc.Client
}

func Dial(url string) (*Client, error) {
	c, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{ec: c}, nil
}

func (c *Client) Call(result interface{}, method string, args ...interface{}) error {
	return c.ec.Call(result, method, args...)
}

func (c *Client) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (ISubscription, error) {
	return c.ec.Subscribe(ctx, namespace, channel, args...)
}
