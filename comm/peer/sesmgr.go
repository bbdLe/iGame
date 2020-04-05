package peer

import (
	"github.com/bbdLe/iGame/comm/session"
	"sync"
	"sync/atomic"
)

type SessionManager interface {
	SessionAccessor

	Add(session.Session)
	Remove(session.Session)
	Count() int

	SetIDBase(base int64)
}

type CoreSessionManager struct {
	sesById sync.Map
	sesIDGen int64
	count int64
}

func (self *CoreSessionManager) SetIDBase(base int64) {
	atomic.StoreInt64(&self.count, base)
}

func (self *CoreSessionManager) Count() int {
	return int(atomic.LoadInt64(&self.count))
}

func (self *CoreSessionManager) Add(sess session.Session) {
	id := atomic.AddInt64(&self.count, 1)
	atomic.AddInt64(&self.count, 1)

	sess.(interface{
		SetID(int64)
	}).SetID(id)

	self.sesById.Store(sess.ID(), sess)
}

func (self *CoreSessionManager) Remove(sess session.Session) {
	atomic.AddInt64(&self.count, -1)
	self.sesById.Delete(sess.ID())
}

func (self *CoreSessionManager) GetSession(id int64) session.Session {
	if v, ok := self.sesById.Load(id); ok {
		return v.(session.Session)
	}

	return nil
}

func (self *CoreSessionManager) VisitSession(cb func(session.Session) bool) {
	self.sesById.Range(func(key, value interface{}) bool {
		return cb(value.(session.Session))
	})
}

func (self *CoreSessionManager) CloseAllSession() {
	self.VisitSession(func(sess session.Session) bool {
		sess.Close()
		return true
	})
}

func (self *CoreSessionManager) SessionCount() int {
	return self.Count()
}