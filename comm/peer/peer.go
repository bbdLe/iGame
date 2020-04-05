package peer

import (
	"github.com/bbdLe/iGame/comm/eventqueue"
	"github.com/bbdLe/iGame/comm/session"
)

type Peer interface {
	Start()

	Stop()

	TypeName() string
}

type Property interface {
	Name() string

	Address() string

	Queue() eventqueue.EventQueue

	SetName(string)

	SetAddr(string)

	SetQueue(eventqueue.EventQueue)
}

type GenericPeer interface {
	Peer
	Property
}

type ContextSet interface {
	SetContext(key interface{}, value interface{})

	GetContext(key interface{}) (value interface{}, exist bool)

	FetchContext(key interface{}, valuePtr interface{}) (bool)
}

type SessionAccessor interface {
	GetSession(int64) session.Session

	VisitSession(func(session.Session) bool)

	SessionCount() int

	CloseAllSession()
}

type ReadyChecker interface {
	IsReady() bool
}

type CaptureIOPanic interface {
	EnableCaptureIOPanic(v bool)

	CaptureIOPanic() bool
}
