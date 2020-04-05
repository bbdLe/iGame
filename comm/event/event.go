package event

import "github.com/bbdLe/iGame/comm/session"

type RecvMsgEvent struct {
	Ses session.Session
	Msg interface{}
}

type ReplyEvent interface {
	Reply(msg interface{})
}

type SendMsgEvent struct {
	Ses session.Session
	Msg interface{}
}

func (self *RecvMsgEvent) Session() session.Session {
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

func (self *SendMsgEvent) Session() session.Session {
	return self.Ses
}

func (self *SendMsgEvent) Message() interface{} {
	return self.Msg
}
