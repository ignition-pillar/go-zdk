package api

import (
	"github.com/zenon-network/go-zenon/protocol"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-wiki/go-zdk/client"
)

type StatsApi struct {
	client client.Client
}

func NewStatsApi(client client.Client) StatsApi {
	return StatsApi{client}
}

func (s StatsApi) OsInfo() (*api.OsInfoResponse, error) {
	var result api.OsInfoResponse
	err := s.client.Call(&result, "stats.osInfo")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s StatsApi) ProcessInfo() (*api.ProcessInfoResponse, error) {
	var result api.ProcessInfoResponse
	err := s.client.Call(&result, "stats.processInfo")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s StatsApi) NetworkInfo() (*api.NetworkInfoResponse, error) {
	var result api.NetworkInfoResponse
	err := s.client.Call(&result, "stats.networkInfo")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s StatsApi) SyncInfo() (*protocol.SyncInfo, error) {
	var result protocol.SyncInfo
	err := s.client.Call(&result, "stats.syncInfo")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
