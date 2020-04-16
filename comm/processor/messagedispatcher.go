package processor

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
	"reflect"
	"sync"
)

type MessageDispatcher struct {
	handlerByType map[reflect.Type][]EventCallback
	handlerGuard sync.RWMutex
	defaultCallback EventCallback
}

func (self *MessageDispatcher) OnEvent(ev Event) {
	pt := reflect.TypeOf(ev.Message())
	if pt == nil {
		return
	}

	self.handlerGuard.RLock()
	handlers, ok := self.handlerByType[pt.Elem()]
	self.handlerGuard.RUnlock()

	if !ok {
		//log.Logger.Error(fmt.Sprintf("MsgType %s  logic not exist, use default logic", pt.String()))
		if self.defaultCallback != nil {
			self.defaultCallback(ev)
		}
		return
	}

	for _, cb := range handlers {
		cb(ev)
	}
}

func (self *MessageDispatcher) Exist(name string) bool {
	m := comm.MessageMetaByFullName(name)
	if m == nil {
		return false
	}

	self.handlerGuard.RLock()
	defer self.handlerGuard.RUnlock()

	handlers, _ := self.handlerByType[m.Type]
	return len(handlers) > 0
}

func (self *MessageDispatcher) RegisterMessage(name string, cb EventCallback) {
	m := comm.MessageMetaByFullName(name)
	if m == nil {
		panic(fmt.Errorf("Can't Find meta : %s", name))
	}

	self.handlerGuard.Lock()
	defer self.handlerGuard.Unlock()

	handlers, _ := self.handlerByType[m.Type]
	handlers = append(handlers, cb)
	self.handlerByType[m.Type] = handlers
}

func (self *MessageDispatcher) SetDefaultCallback(cb EventCallback) {
	self.defaultCallback = cb
}

func NewMessageDispatcher() *MessageDispatcher {
	return &MessageDispatcher{
		handlerByType: make(map[reflect.Type][]EventCallback),
	}
}