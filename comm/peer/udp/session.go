package udp

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/peer"
)

type DataReader interface {
	ReadData() []byte
}

type DataWriter interface {
	WriteData(data []byte)
}

type udpSession struct {
	*peer.CoreProcBundle
	peer.CoreContextSet

	pInterface comm.Peer
}
