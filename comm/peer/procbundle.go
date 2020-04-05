package peer

import (
	"github.com/bbdLe/iGame/comm/err"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/session"
)

var (
	notHandled = err.NewError("ProcBundle: Transmitter is nil")
)

type CoreProcBundle struct{
	transmit processor.MessageTransmitter
	hooker processor.EventHooker
	callback processor.EventCallback
}

func (self *CoreProcBundle) SetTransmitter(mt processor.MessageTransmitter) {
	self.transmit = mt
}

func (self *CoreProcBundle) SetHooker(eh processor.EventHooker) {
	self.hooker = eh
}

func (self *CoreProcBundle) SetCallback(cb processor.EventCallback) {
	self.callback = cb
}

func (self *CoreProcBundle) GetBundle() *CoreProcBundle {
	return self
}

func (self *CoreProcBundle) ReadMessage(sess session.Session) (msg interface{}, err error) {
	if self.transmit != nil {
		return self.transmit.OnRecvMessage(sess)
	} else {
		return nil, notHandled
	}
}

func (self *CoreProcBundle) SendMessage(ev processor.Event) error {
	if self.hooker != nil {
		ev = self.hooker.OnOutboundEvent(ev)
	}

	if self.transmit != nil && ev != nil {
		return self.transmit.OnSendMessage(ev.Session(), ev.Message())
	} else {
		return notHandled
	}
}

func (self *CoreProcBundle) ProcessEvent(ev processor.Event) {
	if self.hooker != nil {
		ev = self.hooker.OnInboundEvent(ev)
	}

	if self.callback != nil && ev != nil {
		self.callback(ev)
	}
}