package proxy

import (
	"github.com/weseliu/alphanet/eventchannel"
	"log"
)

var channel *eventchannel.EventChannel
func Channel() *eventchannel.EventChannel {
	if channel == nil {
		channel = eventchannel.Channel("proxy")
	}
	return channel
}

func Start(address string, path string, remoteAddress string, remotePath string){
	Channel().Start(address, path, remoteAddress, remotePath)
	subscribeEvent()
}

func subscribeEvent(){
	Channel().Subscribe("Game-SendPlayerMessage", onSendPlayerMessage)
}

func onSendPlayerMessage(playerId int64, msg []byte, length int){
	log.Println("onSendPlayerMessage proxy : ", playerId)
}

func OnPlayerEnter(playerId int64){
	Channel().Publish("Proxy-OnPlayerEnter", playerId)
}

func OnPlayerLeave(playerId int64){
	Channel().Publish("Proxy-OnPlayerLeave", playerId)
}

func SendPlayerMessage(playerId int64, msg []byte, length int){
	Channel().Publish("Proxy-OnPlayerMessage", playerId, msg, length)
}
