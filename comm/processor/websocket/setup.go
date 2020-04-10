package websocket

import "github.com/bbdLe/iGame/comm/processor"

func init() {
	processor.RegProcessor("websocket.ltv", func(bundle processor.ProcessorBundle, cb processor.EventCallback, args ...interface{}) {
		bundle.SetTransmitter(new(WebSocketTransmiiter))
		bundle.SetHooker(new(WebSocketHooker))
		bundle.SetCallback(cb)
	})
}
