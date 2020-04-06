package processor

import (
	"github.com/bbdLe/iGame/comm"
)

type Event interface {
	Session() comm.Session

	Message() interface{}
}

type MessageTransmitter interface {
	OnRecvMessage(sess comm.Session) (data interface{}, err error)

	OnSendMessage(sess comm.Session, data interface{}) error
}

type EventHooker interface {
	OnInboundEvent(input Event) (output Event)

	OnOutboundEvent(input Event) (output Event)
}

// 用户回调
type EventCallback func(ev Event)