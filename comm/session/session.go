package session

import (
	"github.com/bbdLe/iGame/comm/meta"
	"github.com/bbdLe/iGame/comm/peer"
)

type Session interface {
	Raw() interface{}

	Peer() peer.Peer

	Send(msg interface{})

	Close()

	ID() int64
}

type RawPacket struct {
	Data []byte
	MsgID int
}

func (self *RawPacket) Message() interface{} {
	if self.MsgID == 0 {
		return struct{}{}
	}

	meta := meta.MessageMetaByID(self.MsgID)
	if meta == nil {
		return struct{}{}
	}

	obj := meta.NewType()
	if err := meta.Codec.Decode(self.Data, obj); err != nil {
		return struct{}{}
	} else {
		return obj
	}
}