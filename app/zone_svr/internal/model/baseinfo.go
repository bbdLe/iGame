package model

import (
	"fmt"

	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
)

type PlayerBaseInfo struct {
	player *Player

	level int
	exp int64
}

func (self *PlayerBaseInfo) Init(p *Player) {
	self.player = p
}

func (self *PlayerBaseInfo) Tick() {
	log.Logger.Debug(fmt.Sprintf("Now Exp : %d", self.exp))
}

func (self *PlayerBaseInfo) Level() int {
	return self.level
}

func (self *PlayerBaseInfo) Exp() int64 {
	return self.exp
}

func (self *PlayerBaseInfo) AddExp(exp int64) {
	self.exp += exp
}

func (self *PlayerBaseInfo) Player() internal.CommPlayer {
	return self.player
}