package peer

import "github.com/bbdLe/iGame/comm/eventqueue"

type CorePeerProperty struct {
	name string
	queue eventqueue.EventQueue
	addr string
}

func (self *CorePeerProperty) Name() string {
	return self.name
}

func (self *CorePeerProperty) Queue() eventqueue.EventQueue {
	return self.queue
}

func (self *CorePeerProperty) Address() string {
	return self.addr
}

func (self *CorePeerProperty) SetName(v string) {
	self.name = v
}

func (self *CorePeerProperty) SetQueue(v eventqueue.EventQueue) {
	self.queue = v
}

func (self *CorePeerProperty) SetAddress(v string) {
	self.addr = v
}
