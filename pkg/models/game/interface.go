package game

import "github.com/gorilla/websocket"

type Interface interface {
	HandleSession(roomId string, conn *websocket.Conn)
}
