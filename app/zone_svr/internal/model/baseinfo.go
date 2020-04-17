package model

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
)

type PlayerBaseInfoImpl struct {
	player *PlayerImpl

	level int
	exp int64
	name string
}

func (self *PlayerBaseInfoImpl) Init(p *PlayerImpl) {
	self.player = p
}

func (self *PlayerBaseInfoImpl) Tick() {
}

func (self *PlayerBaseInfoImpl) Level() int {
	return self.level
}

func (self *PlayerBaseInfoImpl) Exp() int64 {
	return self.exp
}

func (self *PlayerBaseInfoImpl) AddExp(exp int64) {
	self.exp += exp
}

func (self *PlayerBaseInfoImpl) Player() internal.Player {
	return self.player
}

func (self *PlayerBaseInfoImpl) SetName(n string) {
	self.name = n
}

func (self *PlayerBaseInfoImpl) Name() string {
	return self.name
}