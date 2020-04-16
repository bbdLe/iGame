package model

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
)

type PlayerBaseInfo struct {
	player *Player

	level int
	exp int64
	name string
}

func (self *PlayerBaseInfo) Init(p *Player) {
	self.player = p
}

func (self *PlayerBaseInfo) Tick() {
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

func (self *PlayerBaseInfo) SetName(n string) {
	self.name = n
}

func (self *PlayerBaseInfo) Name() string {
	return self.name
}