package zdk

import (
	"github.com/ignition-pillar/go-zdk/api"
	"github.com/ignition-pillar/go-zdk/client"
)

type Zdk struct {
	Ledger    api.LedgerApi
	Stats     api.StatsApi
	Embedded  api.EmbeddedApi
	Subscribe api.SubscribeApi
}

func NewZdk(client client.Client) *Zdk {
	return &Zdk{
		Ledger:    api.NewLedgerApi(client),
		Stats:     api.NewStatsApi(client),
		Embedded:  api.NewEmbeddedApi(client),
		Subscribe: api.NewSubscribeApi(client),
	}
}
