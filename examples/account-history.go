package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-wiki/go-zdk/client"
	"github.com/zenon-wiki/go-zdk/zdk"
)

func main() {

	fmt.Println("Enter your node address: [ws://127.0.0.1:35998]")
	var url string
	fmt.Scanln(&url)
	url = strings.Trim(url, " ")
	if url == "" {
		url = "ws://127.0.0.1:35998"
	}

	rpc, err := client.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to", url)
	z := zdk.Zdk().SetClient(rpc)

	fmt.Println("Enter an address:")
	var addr string
	fmt.Scanln(&addr)

	address := types.ParseAddressPanic(addr)

	var action string
	var counterParty types.Address
	var amount float64 // Note: use decimal for real financial applications!
	var token string

	var height uint64 = 1
	more := true
	for more {
		blocks, err := z.Ledger.GetBlocksByHeight(address, height, 50)
		if err != nil {
			log.Fatal(err)
		}

		for _, b := range blocks.List {

			if b.IsReceiveBlock() {
				paired := b.PairedAccountBlock
				counterParty = paired.Address

				if counterParty == types.ZeroAddress {
					action = "received genesis from"
				} else {
					amount = float64(paired.Amount.Int64()) / math.Pow(10, float64(paired.TokenInfo.Decimals))
					token = paired.TokenInfo.TokenSymbol
					action = fmt.Sprintf("received %.2f %s from", amount, token)
				}

			} else {

				counterParty = b.ToAddress
				if b.Amount.Int64() == 0 {
					action = "interacted with"
				} else {
					amount = float64(b.Amount.Int64()) / math.Pow(10, float64(b.TokenInfo.Decimals))
					token = b.TokenInfo.TokenSymbol
					action = fmt.Sprintf("sent %.2f %s to", amount, token)
				}
			}

			fmt.Printf("%d: At momentum %d, %s %s\n", b.Height, b.MomentumAcknowledged.Height, action, counterParty)
		}

		height += uint64(len(blocks.List))
		more = height <= uint64(blocks.Count)
	}
}
