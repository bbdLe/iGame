package logic

import (
	"github.com/bbdLe/iGame/comm/processor"

	_ "github.com/bbdLe/iGame/proto"
)

var (
	FrontMsgDispatcher *processor.MessageDispatcher
	BackMsgDispatcher *processor.MessageDispatcher
)

func init() {
	FrontMsgDispatcher = processor.NewMessageDispatcher()
	BackMsgDispatcher = processor.NewMessageDispatcher()
	FrontMsgDispatcher.SetDefaultCallback(ZoneDefaultHanlder)
}