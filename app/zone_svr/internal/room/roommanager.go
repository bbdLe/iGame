package room

import (
	"time"
)

type RoomManager struct {
	RoomMap map[int64]*Room

	LastTickTime time.Time
}

func (self *RoomManager) Tick() {
	for _, room := range self.RoomMap {
		room.Tick()
	}
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		RoomMap: make(map[int64]*Room),
		LastTickTime: time.Now(),
	}
}
