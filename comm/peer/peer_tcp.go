package peer

import (
	"github.com/bbdLe/iGame/comm/session"
	"time"
)

type TCPSocketOption interface {
	SetSocketBuffer(readBuffSize, writeBuffSize int, noDelay bool)

	SetMaxPacketSize(maxSize int)

	SetSocketDeadline(read, write time.Duration)
}

type TCPConnector interface {
	GenericPeer
	TCPSocketOption

	SetReconnectDuration(time.Duration)

	ReconnectDuration() time.Duration

	Session() session.Session

	SetSessionManager(raw interface{})

	Port() int
}
