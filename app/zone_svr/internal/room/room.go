package room

import (
	"fmt"
	"github.com/bbdLe/iGame/proto"

	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
)

type RoomImpl struct {
	AoiManager
	Id int64

	playerMap map[int64]internal.Player
}

func (self *RoomImpl) ID() int64 {
	return self.Id
}

func (self *RoomImpl) SetID(id int64) {
	self.Id = id
}

func (self *RoomImpl) Tick() {
}

func (self *RoomImpl) AddPlayer(p internal.Player) {
	log.Logger.Info(fmt.Sprintf("player[%d] enter room[%d]", p.ID(), self.ID()))

	self.playerMap[p.ID()] = p
	p.SetRoom(self)
	self.OnPlayerEnter(p)
}

func (self *RoomImpl) RemovePlayer(p internal.Player) {
	log.Logger.Info(fmt.Sprintf("player[%d] leave room[%d]", p.ID(), self.ID()))

	p.SetRoom(nil)
	delete(self.playerMap, p.ID())
	self.OnPlayerLeave(p)
}

func (self *RoomImpl) VisitPlayer(f func(p internal.Player)) {
	for _, p := range self.playerMap {
		f(p)
	}
}

func (self *RoomImpl) Broadcast(msg interface{}) {
	log.Logger.Debug("broadcast")
	self.VisitPlayer(func(p internal.Player) {
		p.Send(msg)
	})
}

func (self *RoomImpl) OnPlayerEnter(p internal.Player) {
	var msg proto.BroadcastMsgRes
	msg.Msg = fmt.Sprintf("欢迎%s进入房间", p.Name())
	msg.Type = proto.MSG_TYPE_SYSTEM
	self.Broadcast(&msg)
}

func (self *RoomImpl) OnPlayerLeave(p internal.Player) {
	var msg proto.BroadcastMsgRes
	msg.Msg = fmt.Sprintf("玩家%s离开了房间", p.Name())
	msg.Type = proto.MSG_TYPE_SYSTEM
	self.Broadcast(&msg)
}

func NewRoom(id int64) *RoomImpl {
	return &RoomImpl{
		Id : id,
		playerMap : make(map[int64]internal.Player),
	}
}