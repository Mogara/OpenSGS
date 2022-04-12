package game

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/Mogara/OpenSGS/pkg/models/room"
)

type gamer struct {
	rooms map[string]*room.Room
}

func NewGame() Interface {
	return &gamer{
		rooms: make(map[string]*room.Room),
	}
}

func (g *gamer) HandleSession(roomId string, conn *websocket.Conn) {
	memberId := conn.RemoteAddr().String()
	log := log.WithFields(log.Fields{
		"room":   roomId,
		"member": memberId,
	})
	if r, exist := g.rooms[roomId]; exist {
		if r.ExistMember(memberId) {
			log.Warn("member already exist")
		} else {
			if r.AddMember(memberId, conn) {
				log.Info("member join")
			} else {
				_ = conn.WriteMessage(websocket.TextMessage, []byte("{\"message\": \"room is full\"}"))
				log.Warn("room is full")
				conn.Close()
				return
			}
		}
	} else {
		r := room.NewRoom(roomId, 2)
		_ = r.AddMember(memberId, conn)
		log.Info("member join")
		g.rooms[roomId] = r
	}
}
