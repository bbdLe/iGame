package model

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/proto"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
)

type PlayerCmpt interface {
	Tick()

	Init(*PlayerImpl)
}

type PlayerImpl struct {
	Ses comm.Session
	SessionID int64

	Cmpts    []PlayerCmpt
	baseInfo PlayerBaseInfoImpl

	room internal.Room

	HeartBeatTime time.Time
	Status int32
}

func (self *PlayerImpl) RegCmpt(m PlayerCmpt) {
	m.Init(self)
	self.Cmpts = append(self.Cmpts, m)
}

func (self *PlayerImpl) Init() {
	self.RegCmpt(&self.baseInfo)
}

func (self *PlayerImpl) Tick() {
	for _, cmpt := range self.Cmpts {
		cmpt.Tick()
	}

	if self.HeartBeatTime.Add(time.Second * 3).Before(time.Now()) {
		log.Logger.Info(fmt.Sprintf("player[%d] heartbeat time expire", self.SessionID))
		internal.GameMgr.KickPlayer(self)
	}
}

func (self *PlayerImpl) SetHeartBeat(t time.Time) {
	self.HeartBeatTime = t
}

func (self *PlayerImpl) HeartBeat() time.Time {
	return self.HeartBeatTime
}

func (self *PlayerImpl) Session() comm.Session {
	return self.Ses
}

func (self *PlayerImpl) ID() int64 {
	return self.SessionID
}

func (self *PlayerImpl) Name() string {
	return self.baseInfo.Name()
}

func (self *PlayerImpl) Type() internal.EntityType {
	return internal.PlayerType
}

func (self *PlayerImpl) Room() internal.Room {
	return self.room
}

func (self *PlayerImpl) SetRoom(room internal.Room) {
	self.room = room
}

func (self *PlayerImpl) OnLogout() {
	if self.Room() != nil {
		self.Room().RemovePlayer(self)
	}
}

func (self *PlayerImpl) OnLogin() {
	self.baseInfo.SetName(fmt.Sprintf("player_%d", self.ID()))
}

func (self *PlayerImpl) Send(msg interface{}) {
	internal.Send2Player(self, msg)
}

func (self *PlayerImpl) BaseInfo() internal.PlayerBaseInfo {
	return &self.baseInfo
}

func (self *PlayerImpl) EnterView(entity internal.Entity) {
	self.Send(&proto.EnterViewReq{
		EntityId:  entity.ID(),
		EntityType: int32(entity.Type()),
		Pos: entity.Pos(),
	})
}

func (self *PlayerImpl) Pos() *proto.Pos {
	return &proto.Pos{
		X: self.baseInfo.Pos().X(),
		Y: self.baseInfo.Pos().Y(),
	}
}

func (self *PlayerImpl) LeaveView(entity internal.Entity) {
	self.Send(&proto.LeaveViewReq{
		EntityId: entity.ID(),
		EntityType: int32(entity.ID()),
	})
}

func (self *PlayerImpl) SendPos(entity internal.Entity) {
	log.Logger.Debug("send pos")
	self.Send(&proto.PosChangeReq{
		Pos: entity.Pos(),
		EntityId : entity.ID(),
		EntityType: int32(entity.Type()),
	})
}

func (self *PlayerImpl) EnterRoom(roomId int64) int64 {
	self.LeaveRoom()

	room, ok := internal.RoomMgr.GetRoom(roomId)
	if !ok {
		return -1
	}

	room.AddPlayer(self)
	return 0
}

func (self *PlayerImpl) LeaveRoom() {
	if self.Room() != nil {
		self.Room().RemovePlayer(self)
		self.SetRoom(nil)
	}
}

func NewPlayer(sessionID int64, ses comm.Session) *PlayerImpl {
	self := &PlayerImpl{
		SessionID: sessionID,
		Ses : ses,
		HeartBeatTime: time.Now(),
	}
	self.Init()
	return self
}

func init() {
}
