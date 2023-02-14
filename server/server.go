package server

import "github.com/gorilla/websocket"

type WsServer struct {
	Conns map[uint]*websocket.Conn
}

var Server *WsServer

func NewServer() {
	Server = &WsServer{
		Conns: make(map[uint]*websocket.Conn, 0),
	}
}
