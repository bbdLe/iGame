package peer

import "fmt"

type PeerCreateFunc func() Peer

var (
	peerMap = make(map[string]PeerCreateFunc)
)

func RegPeerCreateor(f PeerCreateFunc) {
	t := f()
	peerMap[t.TypeName()] = f
}

func NewPeer(name string) Peer {
	if f, ok := peerMap[name]; ok {
		return f()
	} else {
		panic(fmt.Errorf("New Peer Type failed: %s", name))
	}
}

func GetPeerCreateorList() []string {
	var l []string

	for n := range peerMap {
		l = append(l, n)
	}

	return l
}