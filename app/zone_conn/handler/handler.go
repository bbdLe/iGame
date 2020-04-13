package handler

import "github.com/bbdLe/iGame/comm/processor"

var (
	MsgDispatcher *processor.MessageDispatcher
)

func init() {
	MsgDispatcher = processor.NewMessageDispatcher()
	MsgDispatcher.SetDefaultCallback(ZoneDefaultHanlder)
}