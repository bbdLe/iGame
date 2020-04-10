package websocket

import "github.com/bbdLe/iGame/comm/processor"

type WebSocketHooker struct {}

func (self WebSocketHooker) OnInboundEvent(inputEvent processor.Event) (outputEvent processor.Event) {
	// to do rpc

	return inputEvent
}

func (self WebSocketHooker) OnOutboundEvent(inputEvent processor.Event) (outputEvent processor.Event) {
	return inputEvent
}
