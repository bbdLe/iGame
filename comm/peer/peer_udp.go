package peer

import "github.com/bbdLe/iGame/comm"

type UDPConnector interface {
	comm.GenericPeer

	Session() comm.Session
}
