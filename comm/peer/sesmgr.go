package peer

import (
	"github.com/bbdLe/iGame/comm"
	"sync"
	"sync/atomic"
)

type SessionManager interface {
	comm.SessionAccessor

	Add(comm.Session)
	Remove(comm.Session)
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

func (self *CoreSessionManager) Add(sess comm.Session) {
	id := atomic.AddInt64(&self.count, 1)
	atomic.AddInt64(&self.count, 1)

	sess.(interface{
		SetID(int64)
	}).SetID(id)

	self.sesById.Store(sess.ID(), sess)
}

func (self *CoreSessionManager) Remove(sess comm.Session) {
	atomic.AddInt64(&self.count, -1)
	self.sesById.Delete(sess.ID())
}

func (self *CoreSessionManager) GetSession(id int64) comm.Session {
	if v, ok := self.sesById.Load(id); ok {
		return v.(comm.Session)
	}

	return nil
}

func (self *CoreSessionManager) VisitSession(cb func(comm.Session) bool) {
	self.sesById.Range(func(key, value interface{}) bool {
		return cb(value.(comm.Session))
	})
}

func (self *CoreSessionManager) CloseAllSession() {
	self.VisitSession(func(sess comm.Session) bool {
		sess.Close()
		return true
	})
}

func (self *CoreSessionManager) SessionCount() int {
	return self.Count()
}