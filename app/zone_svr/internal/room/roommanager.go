package room

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"sync/atomic"
	"time"
)

type RoomManager struct {
	RoomMap map[int64]*RoomImpl
	LastTickTime time.Time

	BaseID int64
}

func (self *RoomManager) Tick() {
	if self.LastTickTime.Add(time.Millisecond * 50).Before(time.Now()) {
		return
	}

	for _, room := range self.RoomMap {
		room.Tick()
	}
	self.LastTickTime = time.Now()
}

func (self *RoomManager) NewRoom() internal.Room {
	id := atomic.AddInt64(&self.BaseID, 1)
	r := NewRoom(id)
	self.RoomMap[id] = r
	return r
}

func (self *RoomManager) GetRoom(id int64) (internal.Room, bool) {
	r, ok := self.RoomMap[id]
	return r, ok
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		RoomMap: make(map[int64]*RoomImpl),
		LastTickTime: time.Now(),
	}
}

func init() {
	internal.RoomMgr = NewRoomManager()
}
