package session

import "github.com/bbdLe/iGame/comm/peer"

type Session interface {
	Raw() interface{}

	Peer() peer.Peer

	Send(msg interface{})

	Close()

	ID() int64
}
