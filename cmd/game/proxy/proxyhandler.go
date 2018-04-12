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

var subscribed = false
func subscribeEvent(){
	Channel().Subscribe("Proxy-OnPlayerEnter", onPlayerEnter)
	Channel().Subscribe("Proxy-OnPlayerLeave", onPlayerLeave)
	Channel().Subscribe("Proxy-OnPlayerMessage", onPlayerMessage)
}

func onPlayerEnter(playerId int64){
	log.Println("onPlayerEnter:", playerId)
	SendPlayerMessage(playerId, []byte("game server onPlayerEnter"), 10)
}

func onPlayerLeave(playerId int64){
	log.Println("onPlayerLeave:", playerId)
	SendPlayerMessage(playerId, []byte("game server onPlayerLeave"), 10)
}

func onPlayerMessage(playerId int64, msg []byte, length int){
	log.Println("onPlayerMessage:", playerId)
}

func SendPlayerMessage(playerId int64, msg []byte, length int)  {
	Channel().Publish("Game-SendPlayerMessage", playerId, msg, length)
}
