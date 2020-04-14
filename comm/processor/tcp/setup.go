package tcp

import (
	"github.com/bbdLe/iGame/comm/processor"
)

func init() {
	processor.RegProcessor("tcp.ltv", func(bundle processor.ProcessorBundle, cb processor.EventCallback, args ...interface{}) {
		bundle.SetTransmitter(new(TCPMessageTransmitter))
		bundle.SetHooker(new(MsgHooker))
		bundle.SetCallback(processor.NewQueuedEventCallback(cb))
	})
}
