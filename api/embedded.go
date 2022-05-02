package api

import (
	"github.com/zenon-wiki/go-zdk/api/embedded"
	"github.com/zenon-wiki/go-zdk/client"
)

type EmbeddedApi struct {
	Accelerator embedded.AcceleratorApi
	Pillar      embedded.PillarApi
	Plasma      embedded.PlasmaApi
	Sentinel    embedded.SentinelApi
	Spork       embedded.SporkApi
	Stake       embedded.StakeApi
	Swap        embedded.SwapApi
	Token       embedded.TokenApi
}

func NewEmbeddedApi(client client.Client) EmbeddedApi {
	return EmbeddedApi{
		Accelerator: embedded.NewAcceleratorApi(client),
		Pillar:      embedded.NewPillarApi(client),
		Plasma:      embedded.NewPlasmaApi(client),
		Sentinel:    embedded.NewSentinelApi(client),
		Spork:       embedded.NewSporkApi(client),
		Stake:       embedded.NewStakeApi(client),
		Swap:        embedded.NewSwapApi(client),
		Token:       embedded.NewTokenApi(client),
	}
}
