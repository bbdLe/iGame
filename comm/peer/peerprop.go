package peer

import (
	"github.com/bbdLe/iGame/comm"
)

type CorePeerProperty struct {
	name  string
	queue comm.EventQueue
	addr  string
}

func (self *CorePeerProperty) Name() string {
	return self.name
}

func (self *CorePeerProperty) Queue() comm.EventQueue {
	return self.queue
}

func (self *CorePeerProperty) Address() string {
	return self.addr
}

func (self *CorePeerProperty) SetName(v string) {
	self.name = v
}

func (self *CorePeerProperty) SetQueue(v comm.EventQueue) {
	self.queue = v
}

func (self *CorePeerProperty) SetAddress(v string) {
	self.addr = v
}
