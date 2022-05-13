package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *WsServer) Start(addr string) error {
	router := gin.Default()
	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, nil)
	})
	var handlers []gin.HandlerFunc
	handlers = append(handlers, s.Middlewares...)
	handlers = append(handlers, func(context *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		client := NewClientConnection(context, conn)

		client.onClose = func() {
			if s.onCloseConnection != nil {
				s.onCloseConnection(client)
			}
		}
		client.onMessage = func(message Message) {
			if s.onMessage != nil {
				s.onMessage(message, client)
			}
		}

		go client.read()

		s.onConnection(client)
	})
	router.Any("/ws/socket", handlers...)

	return router.Run(addr)
}

