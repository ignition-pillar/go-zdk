//go:build exclude

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ignition-pillar/go-zdk/client"
	"github.com/ignition-pillar/go-zdk/zdk"
	"github.com/zenon-network/go-zenon/rpc/api/subscribe"
)

func main() {

	fmt.Println("Enter your node address: [ws://127.0.0.1:35998]")
	var url string
	fmt.Scanln(&url)
	url = strings.Trim(url, " ")
	if url == "" {
		url = "ws://127.0.0.1:35998"
	}
	rpc, err := client.NewClient(url)
	if err != nil {
		log.Fatal(err)
	}
	z := zdk.NewZdk(rpc)

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
