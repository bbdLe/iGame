package ZoneMsgDispatcher

import (
	"github.com/bbdLe/iGame/app/zone_svr/player"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/processor"
	"log"
	"sync"
)

type ZoneMsgDispatcher struct {
	handlerByID map[int16]MsgCallBack
	handlerGuard sync.Mutex
}

type MsgCallBack func(player *player.Player, ev processor.Event)

func (self *ZoneMsgDispatcher) Register(msgID int16, cb MsgCallBack) {
	self.handlerGuard.Lock()
	defer self.handlerGuard.Unlock()
	self.handlerByID[msgID] = cb
}

func (self *ZoneMsgDispatcher) OnEvent(player *player.Player, ev processor.Event) {
	log.Println(ev)
	meta := comm.MessageMetaByMsg(ev.Message())
	if meta == nil {
		return
	}

	self.handlerGuard.Lock()
	handler, ok := self.handlerByID[int16(meta.MsgId)]
	self.handlerGuard.Unlock()
	log.Println("====", self.handlerByID, ok, meta.MsgId)

	if ok {
		log.Println("====")
		handler(player, ev)
	}
}

func NewZoneMsgDispather() *ZoneMsgDispatcher {
	return &ZoneMsgDispatcher{
		handlerByID:  make(map[int16]MsgCallBack),
	}
}