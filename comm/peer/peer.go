package peer

type Peer interface {
	Start()

	Stop()

	TypeName() string
}