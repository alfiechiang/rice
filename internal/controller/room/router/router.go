package router

import "rice/internal/controller/room"

type HandlerFunc func(interface{})

var RouterMap map[string]HandlerFunc

type RoomRouter struct {
	RouterMap map[string]HandlerFunc
	Room      *room.GameRoom
}

func (r *RoomRouter) SetRouterMap() {

	router := make(map[string]HandlerFunc, 0)
	router["room.enter"] = r.Room.EnterRoom
	r.RouterMap = router
}

func (r *RoomRouter) Exec(name string,data interface{}) {
	h := r.RouterMap[name]
	if h != nil {
		h(data)
	}
}
