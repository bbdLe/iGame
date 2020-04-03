package processor

import "github.com/bbdLe/iGame/comm/session"

type Event interface {
	Session() session.Session

	Message() interface{}
}

type MessageTransmitter interface {
	OnRecvMessage(sess session.Session) (data interface{}, err error)

	OnSendMessage(sess session.Session, data interface{}) error
}

type EventHooker interface {
	OnInboundEvent(input Event) (output Event)

	OnOutboundEvent(input Event) (output Event)
}

// 用户回调
type EventCallback func(ev Event)