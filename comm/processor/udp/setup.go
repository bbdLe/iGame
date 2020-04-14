package udp

import "github.com/bbdLe/iGame/comm/processor"

func init() {
	processor.RegProcessor("udp.ltv", func(bundle processor.ProcessorBundle, cb processor.EventCallback, args ...interface{}) {
		bundle.SetTransmitter(new(UDPMessageTransmitter))
		bundle.SetHooker(new(MsgHooker))
		bundle.SetCallback(processor.NewQueuedEventCallback(cb))
	})
}
