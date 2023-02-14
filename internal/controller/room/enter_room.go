package room

import (
	"fmt"
	"rice/server"

	"github.com/mitchellh/mapstructure"
)

type EnterRoomService struct {
	PlayerID uint
}

func (r *GameRoom) EnterRoom(data interface{}) {

	var service EnterRoomService
	mapstructure.Decode(data, &service)
	conn:=server.Server.Conns[service.PlayerID]
	player :=Player{Conn: conn}
	r.Players=append(r.Players, player)

	fmt.Println("ssss")

}
