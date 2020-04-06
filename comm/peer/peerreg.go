package peer

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
)

type PeerCreateFunc func() comm.Peer

var (
	peerMap = make(map[string]PeerCreateFunc)
)

func RegPeerCreateor(f PeerCreateFunc) {
	t := f()
	peerMap[t.TypeName()] = f
}

func NewPeer(name string) comm.Peer {
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

func NewGenericPeer(peerType, name, addr string, q comm.EventQueue) comm.GenericPeer {
	p := NewPeer(peerType)
	gp := p.(comm.GenericPeer)
	gp.SetName(name)
	gp.SetAddress(addr)
	gp.SetQueue(q)

	return gp
}