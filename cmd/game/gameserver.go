package main

import (
	"github.com/weseliu/alphanet/codec"
	_ "github.com/weseliu/alphanet/codec/pb"
	"github.com/weseliu/alphanet/util"
	"github.com/weseliu/alphanet/db"
	"github.com/weseliu/alphanet/cmd/game/handler"
	"github.com/weseliu/alphanet/cmd/game/logic"
)

func main() {
	codec.RegisterPbMessageMeta()

	config := util.Configs("./conf/game.json")
	db.Instance().Open(config.String("dsn"))

	logic.Start()

	handler.ChannelStart()
	handler.ChannelWait()
}