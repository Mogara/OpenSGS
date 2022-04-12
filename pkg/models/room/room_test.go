package room

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomCapacityLimit(t *testing.T) {
	room := NewRoom("test", 2)
	assert.Equal(t, true, room.AddMember("1", nil))
	assert.Equal(t, true, room.AddMember("2", nil))
	assert.Equal(t, false, room.AddMember("3", nil))
	room.RemoveMember("2")
	assert.Equal(t, true, room.AddMember("3", nil))
}
