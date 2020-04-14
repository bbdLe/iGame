package websocket

import "github.com/bbdLe/iGame/comm/processor"

func init() {
	processor.RegProcessor("ws.ltv", func(bundle processor.ProcessorBundle, cb processor.EventCallback, args ...interface{}) {
		bundle.SetTransmitter(new(WebSocketTransmiiter))
		bundle.SetHooker(new(WebSocketHooker))
		bundle.SetCallback(processor.NewQueuedEventCallback(cb))
	})
}
