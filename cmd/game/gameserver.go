package main

import (
	"github.com/weseliu/alphanet/codec"
	_ "github.com/weseliu/alphanet/codec/pb"
	"github.com/weseliu/alphanet/channel"
	"log"
	"github.com/weseliu/alphanet/util"
	"github.com/weseliu/alphanet/db"
	"time"
)

func main() {
	codec.RegisterPbMessageMeta()

	config := util.Configs("./conf/game.json")
	db.Instance().Open(config.String("dsn"))

	channelClient := channel.NewChannel(config.String("channel_address"),
		time.Duration(config.Int64("channel_timeout")),
		int(config.Int64("channel_input_chan_size")),
		int(config.Int64("channel_output_chan_size")))
	channelClient.Connect()
	channelClient.ReadLoop(func(bytes []byte, param int64) {
		log.Println("Client : ", bytes, "roleId : ", param)
		//channelClient.Send([]byte("BBBBBBBBBBBBB"), nil, param)
	})
}