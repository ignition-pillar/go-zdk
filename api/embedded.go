package api

import (
	"github.com/zenon-wiki/go-zdk/api/embedded"
	"github.com/zenon-wiki/go-zdk/client"
)

type EmbeddedApi struct {
	client client.IClient

	//Accelerator embedded.AcceleratorApi
	Pillar   embedded.PillarApi
	Plasma   embedded.PlasmaApi
	Sentinel embedded.SentinelApi
	Spork    embedded.SporkApi
	Stake    embedded.StakeApi
	Swap     embedded.SwapApi
	Token    embedded.TokenApi
}

func (e *EmbeddedApi) SetClient(client client.IClient) {
	e.client = client

	//e.Accelerator.SetClient(client)
	e.Pillar.SetClient(client)
	e.Plasma.SetClient(client)
	e.Sentinel.SetClient(client)
	e.Spork.SetClient(client)
	e.Stake.SetClient(client)
	e.Swap.SetClient(client)
	e.Token.SetClient(client)
}

func NewEmbeddedApi() EmbeddedApi {
	e := EmbeddedApi{
		//Accelerator: embedded.AcceleratorApi{},
		Pillar:   embedded.PillarApi{},
		Plasma:   embedded.PlasmaApi{},
		Sentinel: embedded.SentinelApi{},
		Spork:    embedded.SporkApi{},
		Stake:    embedded.StakeApi{},
		Swap:     embedded.SwapApi{},
		Token:    embedded.TokenApi{},
	}
	return e
}
