package room

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	*sync.Mutex
	id       string
	capacity int
	currnet  int
	member   map[string]*websocket.Conn
}

func NewRoom(id string, capacity int) *Room {
	return &Room{
		id:       id,
		Mutex:    &sync.Mutex{},
		capacity: capacity,
		currnet:  0,
		member:   make(map[string]*websocket.Conn),
	}
}

func (r *Room) AddMember(id string, conn *websocket.Conn) bool {
	r.Lock()
	defer r.Unlock()
	if r.currnet >= r.capacity {
		return false
	}
	r.member[id] = conn
	r.currnet++
	return true
}

func (r *Room) RemoveMember(id string) {
	r.Lock()
	defer r.Unlock()
	delete(r.member, id)
	r.currnet--
}

func (r *Room) ExistMember(id string) bool {
	r.Lock()
	defer r.Unlock()
	_, exist := r.member[id]
	return exist
}
