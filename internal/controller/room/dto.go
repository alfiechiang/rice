package room

import (
	"github.com/gorilla/websocket"
)

type WsReq struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type WsResp struct {
	Name string
	Data interface{}
}

type WsServer struct {
	Conn    *websocket.Conn
	Rooms   []GameRoom
	Req     *WsReq
	OutChan chan *WsResp
}

type Player struct {
	Conn *websocket.Conn
}

type GameRoom struct {
	RoomID  int64 `mapstructure:"roomID"`
	Players []Player
}

var Rooms = make([]GameRoom, 0)
