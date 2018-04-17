package handler

import (
	"github.com/weseliu/alphanet/channel"
	"github.com/weseliu/alphanet/util"
	"time"
	"log"
	"github.com/weseliu/alphanet/cmd/global/encoder"
	"github.com/weseliu/alphanet/core"
)

var proxyChannel *channel.Channel

func ChannelStart() {
	config := util.Configs("./conf/game.json")
	proxyChannel = channel.NewChannel(config.String("channel_address"),
		time.Duration(config.Int64("channel_timeout")),
		int(config.Int64("channel_input_chan_size")),
		int(config.Int64("channel_output_chan_size")))
	proxyChannel.Connect()
}

func ChannelWait()  {
	proxyChannel.ReadLoop(onGameMessage)
}

func onGameMessage(bytes []byte, param int64) {
	log.Print("onGameMessage : ", string(bytes))
	var msg = encoder.DecodeMsg(bytes)
	if msg != nil {
		core.MsgCenter().DispatchMessage(msg, param)
	}
}

func SendMessage(roldId int64, msg interface{}, code int32)  {
	data := encoder.EncodeMsg(code, msg)
	if data != nil {
		proxyChannel.Send(data, nil, roldId)
	}
}