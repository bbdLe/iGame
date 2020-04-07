package udp

import (
	"github.com/bbdLe/iGame/comm"
)

type UDPMessageTransmitter struct {
}

func (UDPMessageTransmitter) OnRecvMessage(sess comm.Session) (msg interface{}, err error) {
	return nil, nil
}

func (UDPMessageTransmitter) OnSendMessage(sess comm.Session, msg interface{}) (err error) {
	return nil
}
