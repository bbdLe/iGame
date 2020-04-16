package room

import (
	"fmt"
	"github.com/bbdLe/iGame/proto"

	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
)

type Room struct {
	Id int64

	playerMap map[int64]internal.CommPlayer
}

func (self *Room) ID() int64 {
	return self.Id
}

func (self *Room) SetID(id int64) {
	self.Id = id
}

func (self *Room) Tick() {
	for _, p := range self.playerMap {
		p.BaseInfo().AddExp(10)
	}
}

func (self *Room) AddPlayer(p internal.CommPlayer) {
	log.Logger.Info(fmt.Sprintf("player[%d] enter room[%d]", p.ID(), self.ID()))

	self.playerMap[p.ID()] = p
	p.SetRoom(self)
	self.OnPlayerEnter(p)
}

func (self *Room) RemovePlayer(p internal.CommPlayer) {
	log.Logger.Info(fmt.Sprintf("player[%d] leave room[%d]", p.ID(), self.ID()))

	p.SetRoom(nil)
	delete(self.playerMap, p.ID())
}

func (self *Room) VisitPlayer(f func(p internal.CommPlayer)) {
	for _, p := range self.playerMap {
		f(p)
	}
}

func (self *Room) Broadcast(msg interface{}) {
	log.Logger.Debug("broadcast")
	self.VisitPlayer(func(p internal.CommPlayer) {
		p.Send(msg)
	})
}

func (self *Room) OnPlayerEnter(p internal.CommPlayer) {
	var msg proto.BroadcastMsgRes
	msg.Msg = fmt.Sprintf("欢迎%s进入游戏", p.Name())
	msg.Type = proto.MSG_TYPE_SYSTEM
	self.Broadcast(&msg)
}

func NewRoom(id int64) *Room {
	return &Room{
		Id : id,
		playerMap : make(map[int64]internal.CommPlayer),
	}
}