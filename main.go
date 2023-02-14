package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rice/internal/controller/room"
	roomRouter "rice/internal/controller/room/router"
	"rice/router"
	"rice/server"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}




func WsHandler(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("playerID"))
	playerID := uint(ID)
	if err != nil {
		c.JSON(500, "解析玩家有誤")
		return
	}
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	server.Server.Conns[playerID]=ws
	if err != nil {
		return
	}

	defer ws.Close()

	for {
		//读取ws中的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		body := &room.WsReq{}
		json.Unmarshal(message, body)
		if err := json.Unmarshal(message, body); err != nil {
			log.Panicln("非法格式:", err)
		}

		router.SetRouterMap()
		router.Exec(body.Name)

		for _, v := range room.Rooms {
			r1 := &room.GameRoom{}
			mapstructure.Decode(body.Data, r1)
			if v.RoomID == r1.RoomID {
				router := roomRouter.RoomRouter{
					Room: &v,
				}
				router.SetRouterMap()
				router.Exec(body.Name, body.Data)
			}
		}

		fmt.Println(room.Rooms)

		err = ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}

func main() {
	bindAddress := "localhost:2301"
	server.NewServer()
	r := gin.Default()
	r.GET("/ws/:playerID", WsHandler)
	r.Run(bindAddress)
}
