package model

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
)

type PlayerCmpt interface {
	Tick()

	Init(*Player)
}

type Player struct {
	Ses comm.Session
	SessionID int64

	Cmpts []PlayerCmpt
	baseInfo PlayerBaseInfo

	room internal.CommRoom

	HeartBeatTime time.Time
	Status int32
}

func (self *Player) RegCmpt(m PlayerCmpt) {
	m.Init(self)
	self.Cmpts = append(self.Cmpts, m)
}

func (self *Player) Init() {
	self.RegCmpt(&self.baseInfo)
}

func (self *Player) Tick() {
	for _, cmpt := range self.Cmpts {
		cmpt.Tick()
	}

	if self.HeartBeatTime.Add(time.Second * 3).Before(time.Now()) {
		log.Logger.Info(fmt.Sprintf("player[%d] heartbeat time expire", self.SessionID))
		internal.GameMgr.KickPlayer(self)
	}
}

func (self *Player) SetHeartBeat(t time.Time) {
	self.HeartBeatTime = t
}

func (self *Player) HeartBeat() time.Time {
	return self.HeartBeatTime
}

func (self *Player) Session() comm.Session {
	return self.Ses
}

func (self *Player) ID() int64 {
	return self.SessionID
}

func (self *Player) Name() string {
	return self.baseInfo.Name()
}

func (self *Player) Room() internal.CommRoom {
	return self.room
}

func (self *Player) SetRoom(room internal.CommRoom) {
	self.room = room
}

func (self *Player) OnLogout() {
	if self.Room() != nil {
		self.Room().RemovePlayer(self)
	}
}

func (self *Player) OnLogin() {
	self.baseInfo.SetName(fmt.Sprintf("player_%d", self.ID()))
}

func (self *Player) Send(msg interface{}) {
	internal.Send2Player(self, msg)
}

func (self *Player) BaseInfo() internal.CommPlayerBaseInfo {
	return &self.baseInfo
}

func NewPlayer(sessionID int64, ses comm.Session) *Player {
	self := &Player{
		SessionID: sessionID,
		Ses : ses,
		HeartBeatTime: time.Now(),
	}
	self.Init()
	return self
}

func init() {
}
