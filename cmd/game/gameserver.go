package main

import (
	//"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/channel"
	"log"
	"time"
)

func main() {
	//codec.RegisterPbMessageMeta()
	//config := util.Configs("./conf/connect.json")
	channelSvr := channel.NewChannel(":2011", 10, 100, 100)
	channelSvr.Listener()


	go channelSvr.ReadLoop(func(bytes []byte) {
		log.Println("Server : ", bytes)
	})

	channelClient := channel.NewChannel(":2011", 10, 100, 100)
	channelClient.Connect()

	go func() {
		for {
			channelSvr.Send([]byte("aaaaaaaaaaaa"))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			channelClient.Send([]byte("BBBBBBBBBBBBB"))
			time.Sleep(1 * time.Second)
		}
	}()
	channelSvr.ReadLoop(func(bytes []byte) {
		log.Println("Client : ", bytes)
	})
}