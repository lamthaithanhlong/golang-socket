package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 64 * 1024	// 64KB
)

<<<<<<< Updated upstream
=======
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type ClientConnection struct {
	context *gin.Context
	conn           *websocket.Conn
	onClose        func()
	onMessage      func(message Message)
	data map[string]interface{}
	mu *sync.Mutex
}

func NewClientConnection(context *gin.Context, conn *websocket.Conn) *ClientConnection {
	return &ClientConnection{
		context: context,
		conn: conn,
		mu: &sync.Mutex{},
	}
}
>>>>>>> Stashed changes
