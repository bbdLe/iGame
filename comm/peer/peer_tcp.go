package peer

import (
	"github.com/bbdLe/iGame/comm"
	"time"
)

type TCPSocketOption interface {
	SetSocketBuffer(readBuffSize, writeBuffSize int, noDelay bool)

	SetMaxPacketSize(maxSize int)

	SetSocketDeadline(read, write time.Duration)
}

type TCPConnector interface {
	comm.GenericPeer
	TCPSocketOption

	SetReconnectDuration(time.Duration)

	ReconnectDuration() time.Duration

	Session() comm.Session

	SetSessionManager(raw interface{})

	Port() int
}
