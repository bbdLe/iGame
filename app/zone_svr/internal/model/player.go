package model

import (
	"github.com/bbdLe/iGame/comm"
	"sync"
)

type PlayerCmpt interface {
	Tick()

	Init()
}

var (
	playerMap map[int64]*Player
	playerMapGuard sync.RWMutex
)

type Player struct {
	baseInfo PlayerBaseInfo

	Cmpts []PlayerCmpt
	SessionID int64
	Status int32
	Ses comm.Session
}

func (self *Player) RegCmpt(m PlayerCmpt) {
	self.Cmpts = append(self.Cmpts, m)
}

func (self *Player) Init() {
	self.RegCmpt(&self.baseInfo)
}

func (self *Player) Tick() {
	for _, cmpt := range self.Cmpts {
		cmpt.Tick()
	}
}

func (self *Player) OnLogout() {

}

func NewPlayer(sessionID int64, ses comm.Session) *Player {
	self := &Player{
		SessionID: sessionID,
		Ses : ses,
	}
	self.Init()
	return self
}

func SetPlayer(sessionID int64, player *Player) {
	playerMapGuard.Lock()
	defer playerMapGuard.Unlock()

	playerMap[sessionID] = player
}

func GetPlayer(sessionID int64) (*Player, bool) {
	playerMapGuard.RLock()
	defer playerMapGuard.RUnlock()

	p, ok := playerMap[sessionID]
	return p, ok
}

func DelPlayer(sessionID int64) {
	playerMapGuard.Lock()
	defer playerMapGuard.Unlock()

	delete(playerMap, sessionID)
}

func init() {
	playerMap = make(map[int64]*Player)
}
