package router

import "rice/internal/controller/room"

type HandlerFunc func()

var RouterMap map[string]HandlerFunc

func SetRouterMap() {

	router := make(map[string]HandlerFunc, 0)
	router["room.create"] = room.CreateRoom


	RouterMap = router

}

func Exec(name string) {
	h := RouterMap[name]
	if h != nil {
		h()
	}
}
