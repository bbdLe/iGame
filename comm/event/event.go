package event

import (
	"github.com/bbdLe/iGame/comm"
)

type RecvMsgEvent struct {
	Ses comm.Session
	Msg interface{}
}

type ReplyEvent interface {
	Reply(msg interface{})
}

type SendMsgEvent struct {
	Ses comm.Session
	Msg interface{}
}

func (self *RecvMsgEvent) Session() comm.Session {
	return self.Ses
}

func (self *RecvMsgEvent) Message() interface{} {
	return self.Msg
}

func (self *RecvMsgEvent) Send(msg interface{}) {
	self.Session().Send(msg)
}

func (self *RecvMsgEvent) Reply(msg interface{}) {
	self.Session().Send(msg)
}

func (self *SendMsgEvent) Session() comm.Session {
	return self.Ses
}

func (self *SendMsgEvent) Message() interface{} {
	return self.Msg
}
