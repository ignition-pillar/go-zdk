package embedded

import (
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-wiki/go-zdk/client"
)

type SporkApi struct {
	client client.IClient
}

func (s *SporkApi) SetClient(client client.IClient) {
	s.client = client
}

func (s *SporkApi) GetAll(pageIndex, pageSize uint32) (*embedded.SporkList, error) {
	if pageSize > api.RpcMaxPageSize {
		pageSize = api.RpcMaxPageSize
	}
	var result embedded.SporkList
	err := s.client.Call(&result, "embedded.spork.getAll", pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
