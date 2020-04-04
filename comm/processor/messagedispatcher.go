package processor

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/meta"
	"reflect"
	"sync"
)

type MessageDispatcher struct {
	handlerByType map[reflect.Type][]EventCallback
	handlerGuard sync.RWMutex
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
		return
	}

	for _, cb := range handlers {
		cb(ev)
	}
}

func (self *MessageDispatcher) Exist(name string) bool {
	m := meta.MessageMetaByFullName(name)
	if m == nil {
		return false
	}

	self.handlerGuard.RLock()
	defer self.handlerGuard.RUnlock()

	handlers, _ := self.handlerByType[m.Type]
	return len(handlers) > 0
}

func (self *MessageDispatcher) RegisterMessage(name string, cb EventCallback) {
	m := meta.MessageMetaByFullName(name)
	if m == nil {
		panic(fmt.Errorf("Can't Find meta : %s", name))
	}

	self.handlerGuard.Lock()
	defer self.handlerGuard.Unlock()

	handlers, _ := self.handlerByType[m.Type]
	handlers = append(handlers, cb)
	self.handlerByType[m.Type] = handlers
}

func NewMessageDispatcher() *MessageDispatcher {
	return &MessageDispatcher{
		handlerByType: make(map[reflect.Type][]EventCallback),
	}
}