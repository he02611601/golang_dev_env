package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		// 升級成 websocket 連線
		ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatalln(err)
		}

		// 結束時關閉連線
		defer ws.Close()

		for {
			// 讀取client端發送的訊息，如果沒有訊息會阻塞
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("read error")
				fmt.Println(err)
				break
			}

			if string(message) == "ping" {
				message = []byte("pong")
			}

			// 回傳client端訊息回去
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
