package ZoneMsgDispatcher

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/processor"
	"sync"
)

type ZoneMsgDispatcher struct {
	handlerByID map[int16]MsgCallBack
	handlerGuard sync.Mutex
}

type MsgCallBack func(player internal.CommPlayer, ev processor.Event)

func (self *ZoneMsgDispatcher) Register(msgID int16, cb MsgCallBack) {
	self.handlerGuard.Lock()
	defer self.handlerGuard.Unlock()
	self.handlerByID[msgID] = cb
}

func (self *ZoneMsgDispatcher) OnEvent(player internal.CommPlayer, ev processor.Event) {
	meta := comm.MessageMetaByMsg(ev.Message())
	if meta == nil {
		return
	}

	self.handlerGuard.Lock()
	handler, ok := self.handlerByID[int16(meta.MsgId)]
	self.handlerGuard.Unlock()

	if ok {
		handler(player, ev)
	}
}

func NewZoneMsgDispather() *ZoneMsgDispatcher {
	return &ZoneMsgDispatcher{
		handlerByID:  make(map[int16]MsgCallBack),
	}
}