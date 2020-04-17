package room

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
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

	var freeRoomList []int64
	for _, room := range self.RoomMap {
		room.Tick()

		if room.CanFree(){
			freeRoomList = append(freeRoomList, room.ID())
		}
	}

	for _, id := range freeRoomList {
		self.FreeRoom(id)
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

func (self *RoomManager) FreeRoom(id int64) {
	log.Logger.Debug(fmt.Sprintf("free room[%d]", id))
	delete(self.RoomMap, id)
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
