package room

import (
	"sync"
	"time"
)

func CreateRoom() {
	players := make([]Player, 0)

	var mutex sync.Mutex
	mutex.Lock()
	time.Sleep(1 * time.Second)
	room := GameRoom{
		RoomID:  time.Now().Unix(),
		Players: players,
	}
	Rooms = append(Rooms, room)
	mutex.Unlock()

}
