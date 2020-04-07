package udp

import "github.com/bbdLe/iGame/comm/processor"

type MsgHooker struct {}

func (self MsgHooker) OnInboundEvent(inputEvent processor.Event) (outputEvent processor.Event) {
	// to do rpc

	return inputEvent
}

func (self MsgHooker) OnOutboundEvent(inputEvent processor.Event) (outputEvent processor.Event) {
	return inputEvent
}
