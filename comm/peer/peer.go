package peer

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