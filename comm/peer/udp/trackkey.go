package udp

import (
	"encoding/binary"
	"net"
)

type connTrackKey struct {
	IPHigh uint64
	IPLow uint64
	Port int
}

func newConnTrackKey(addr *net.UDPAddr) *connTrackKey {
	if len(addr.IP) == net.IPv4len {
		return &connTrackKey{
			IPHigh: 0,
			IPLow: uint64(binary.BigEndian.Uint32(addr.IP)),
			Port: addr.Port,
		}
	} else {
		return &connTrackKey{
			IPHigh: uint64(binary.BigEndian.Uint64(addr.IP[:8])),
			IPLow: uint64(binary.BigEndian.Uint64(addr.IP[8:])),
			Port: addr.Port,
		}
	}
}