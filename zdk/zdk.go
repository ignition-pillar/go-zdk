package zdk

import (
	"sync"

	"github.com/zenon-wiki/go-zdk/api"
	"github.com/zenon-wiki/go-zdk/client"
)

var (
	once sync.Once
	nom  *ZDK
)

type ZDK struct {
	Client client.IClient

	Ledger    api.LedgerApi
	Stats     api.StatsApi
	Embedded  api.EmbeddedApi
	Subscribe api.SubscribeApi
}

func (z *ZDK) SetClient(client client.IClient) *ZDK {
	z.Client = client
	z.Ledger.SetClient(client)
	z.Stats.SetClient(client)
	z.Embedded.SetClient(client)
	z.Subscribe.SetClient(client)
	return z
}

func Zdk() *ZDK {
	once.Do(func() {
		nom = &ZDK{
			Ledger:    api.LedgerApi{},
			Stats:     api.StatsApi{},
			Embedded:  api.NewEmbeddedApi(),
			Subscribe: api.SubscribeApi{},
		}
	})
	return nom
}
