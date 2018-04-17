package handler

import (
	"github.com/weseliu/alphanet/channel"
	"github.com/weseliu/alphanet/util"
	"time"
	"github.com/weseliu/alphanet/net"
	"log"
	"github.com/weseliu/alphanet/cmd/global/protocal/connect"
	"github.com/weseliu/alphanet/cmd/global/encoder"
)

var proxyChannel *channel.Channel
var roleSessions map[int64]net.Session

func ChannelStart() {
	roleSessions = make(map[int64]net.Session)

	config := util.Configs("./conf/connect.json")
	proxyChannel = channel.NewChannel(config.String("channel_address"),
		time.Duration(config.Int64("channel_timeout")),
		int(config.Int64("channel_input_chan_size")),
		int(config.Int64("channel_output_chan_size")))
	proxyChannel.Listener()

	go proxyChannel.ReadLoop(fireRemoteMessage)
}

func OnRoleEnter(roleId int64, session net.Session)  {
	roleSessions[roleId] = session
}

func OnRoleExit(session net.Session) {
	for id, v := range roleSessions {
		if v == session {
			delete(roleSessions, id)
		}
	}
}

func fireRemoteMessage(bytes []byte, param int64) {
	log.Println("fireRemoteMessage : ", bytes, "param :", param)
	roleId := param
	session := roleSessions[roleId]
	if session != nil {
		logic := &connect.CMD_LOGIC_SC{
			LogicPkg : bytes,
		}
		session.Send(encoder.EncodeCmd(logic))
	}
}

func SendRemoteMessage(roleId int64, msg []byte) {
	proxyChannel.Send(msg, nil, roleId)
}
