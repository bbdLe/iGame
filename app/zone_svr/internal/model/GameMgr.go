package model

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/proto"
	"time"
)

type GameManager struct {
	lastTickTime time.Time
	tickInterVal time.Duration

	playerMap map[int64]internal.CommPlayer
}

func (self *GameManager) Init() {
	self.lastTickTime = time.Now()
}

func (self *GameManager) Tick() {
	if self.lastTickTime.Add(self.tickInterVal).After(time.Now()) {
		internal.ZoneEventQueue.Post(func() {
			self.TickBase()
			internal.RoomMgr.Tick()
		})
		self.lastTickTime = time.Now()
	}
}

func (self *GameManager) TickBase() {
	self.VisitPlayer(func(p internal.CommPlayer) {
		p.Tick()
	})
}

func (self *GameManager) Start() {
	for {
		self.Tick()

		time.Sleep(time.Millisecond * 20)
	}
}

func (self *GameManager) Stop() {

}

func (self *GameManager) SetPlayer(sessionID int64, player internal.CommPlayer) {
	self.playerMap[sessionID] = player
}

func (self *GameManager) GetPlayer(sessionID int64) (internal.CommPlayer, bool) {
	p, ok := self.playerMap[sessionID]
	return p, ok
}

func (self *GameManager) DelPlayer(sessionID int64) {
	delete(self.playerMap, sessionID)
}

func (self *GameManager) VisitPlayer(f func(player internal.CommPlayer)) {
	for _, p := range self.playerMap {
		f(p)
	}
}

func (self *GameManager) KickPlayer(p internal.CommPlayer) {
	p.OnLogout()
	p.Session().Send(&proto.KickConnReq{
		ClientId: p.ID(),
	})
	self.DelPlayer(p.ID())
}

func NewGameManager() *GameManager {
	return &GameManager{
		tickInterVal: time.Millisecond * 50,
		lastTickTime: time.Now(),
		playerMap : make(map[int64]internal.CommPlayer),
	}
}

func init() {
	internal.GameMgr = NewGameManager()
}