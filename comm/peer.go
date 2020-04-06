package comm

type Peer interface {
	Start()

	Stop()

	TypeName() string
}

type Property interface {
	Name() string

	Address() string

	Queue() EventQueue

	SetName(string)

	SetAddress(string)

	SetQueue(EventQueue)
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
	GetSession(int64) Session

	VisitSession(func(Session) bool)

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
