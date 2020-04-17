package model

import (
	"fmt"
	"time"

	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/proto"
)

type GameManagerImpl struct {
	lastTickTime time.Time
	tickInterVal time.Duration

	playerMap map[int64]internal.Player
}

func (self *GameManagerImpl) Init() {
	self.lastTickTime = time.Now()
}

func (self *GameManagerImpl) Tick() {
	if self.lastTickTime.Add(self.tickInterVal).After(time.Now()) {
		internal.ZoneEventQueue.Post(func() {
			self.TickBase()
			internal.RoomMgr.Tick()
		})
		self.lastTickTime = time.Now()
	}
}

func (self *GameManagerImpl) TickBase() {
	self.VisitPlayer(func(p internal.Player) {
		p.Tick()
	})
}

func (self *GameManagerImpl) Start() {
	for {
		self.Tick()

		time.Sleep(time.Millisecond * 20)
	}
}

func (self *GameManagerImpl) Stop() {

}

func (self *GameManagerImpl) SetPlayer(sessionID int64, player internal.Player) {
	self.playerMap[sessionID] = player
}

func (self *GameManagerImpl) GetPlayer(sessionID int64) (internal.Player, bool) {
	p, ok := self.playerMap[sessionID]
	return p, ok
}

func (self *GameManagerImpl) DelPlayer(sessionID int64) {
	delete(self.playerMap, sessionID)
}

func (self *GameManagerImpl) VisitPlayer(f func(player internal.Player)) {
	for _, p := range self.playerMap {
		f(p)
	}
}

func (self *GameManagerImpl) KickPlayer(p internal.Player) {
	p.OnLogout()
	p.Session().Send(&proto.KickConnReq{
		ClientId: p.ID(),
	})
	self.DelPlayer(p.ID())
}

func (self *GameManagerImpl) OnPlayerLogin(p internal.Player) {
	self.VisitPlayer(func(p internal.Player) {
		p.Send(&proto.BroadcastMsgRes{
			Msg:  fmt.Sprintf("欢迎%s登陆游戏", p.Name()),
			Type: proto.MSG_TYPE_SYSTEM,
		})
	})
}

func NewGameManager() *GameManagerImpl {
	return &GameManagerImpl{
		tickInterVal: time.Millisecond * 50,
		lastTickTime: time.Now(),
		playerMap : make(map[int64]internal.Player),
	}
}

func init() {
	internal.GameMgr = NewGameManager()
}