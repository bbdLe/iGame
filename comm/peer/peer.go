package peer

import "github.com/bbdLe/iGame/comm/session"

type Peer interface {
	Start()

	Stop()

	TypeName() string
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