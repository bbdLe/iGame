package peer

import (
	"time"

	"github.com/bbdLe/iGame/comm"
)

type WSAcceptor interface {
	comm.GenericPeer

	comm.SessionAccessor

	SetHttps(certfile, keyfile string)

	SetUpgrader(upgrader interface{})

	Port() int
}

type WSConnector interface {
	comm.GenericPeer

	SetReconnectDuration(time.Duration)

	ReconnectDuration() time.Duration

	Session() comm.Session

	SetSessionManager(raw interface{})

	Port() int
}
