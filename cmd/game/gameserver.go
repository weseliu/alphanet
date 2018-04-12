package main

import (
	//"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/channel"
	"log"
)

func main() {
	//codec.RegisterPbMessageMeta()
	//config := util.Configs("./conf/connect.json")

	channelClient := channel.NewChannel(":2011", 10, 1, 1)
	channelClient.Connect()

	go func() {
		for {
			channelClient.Send([]byte("BBBBBBBBBBBBB"), nil)
		}
	}()

	channelClient.ReadLoop(func(bytes []byte) {
		log.Println("Client : ", bytes)
	})
}