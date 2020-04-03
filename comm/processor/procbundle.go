package processor

import "github.com/bbdLe/iGame/comm/eventqueue"

type ProcessorBundle interface {
	SetTransmitter(mt MessageTransmitter)

	SetHooker(eh EventHooker)

	SetCallback(cb EventCallback)
}

func NewQueueEventCallback(cb EventCallback) EventCallback {
	return func(ev Event) {
		if ev == nil {
			return
		}

		eventqueue.SessionCall(ev.Session(), func() {
			cb(ev)
		})
	}
}

type MultHooker []EventHooker

func (self MultHooker) OnInboundEvent(input Event) Event {
	for _, h := range self {
		input = h.OnInboundEvent(input)
		if input == nil {
			break
		}
	}

	return input
}

func (self MultHooker) OnOutboundEvent(input Event) Event {
	for _, h := range self {
		input = h.OnOutboundEvent(input)
		if input == nil {
			break
		}
	}

	return input
}