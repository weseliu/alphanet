package handler

import (
	"github.com/weseliu/alphanet/channel"
	"github.com/weseliu/alphanet/util"
	"time"
	"github.com/weseliu/alphanet/net"
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
	SendRemoteMessage(roleId, []byte("player enter!"))
}

func OnRoleExit(session net.Session) {
	for id, v := range roleSessions {
		if v == session {
			delete(roleSessions, id)
		}
	}
}

func fireRemoteMessage(bytes []byte, param int64) {
	roleId := param
	session := roleSessions[roleId]
	if session != nil {
		session.Send(bytes)
	}
}

func SendRemoteMessage(roleId int64, msg []byte) {
	proxyChannel.Send(msg, nil, roleId)
}
