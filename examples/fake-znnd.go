package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zenon-network/go-zenon/rpc/api/subscribe"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/zdk"
)

func main() {

	url := "ws://127.0.0.1:35998"
	rpc, err := client.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	z := zdk.Zdk().SetClient(rpc)

	ms := make(chan []subscribe.Momentum)

	sub, err := z.Subscribe.ToMomentums(ms)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case m := <-ms:
			detailed, err := z.Ledger.GetDetailedMomentumsByHeight(m[0].Height, 1)
			if err != nil {
				log.Fatal(err)
			}
			dM := detailed.List[0].Momentum
			numTx := len(detailed.List[0].AccountBlocks)
			fmt.Printf("[Momentum inserted] Height: %d, Hash: %s, Timestamp: %d, Pillar producer address: %s, Current time: %s, Txs: %d\n", m[0].Height, m[0].Hash, dM.TimestampUnix, dM.Producer, time.Now().Format("2006-01-02 15:04:05"), numTx)
		}
	}
	sub.Unsubscribe()
}
