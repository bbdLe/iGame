package peer

import (
	"sync"
	"sync/atomic"
)

type CoreRunningTag struct {
	running int64
	stoppingWg sync.WaitGroup
	stopping int64
}

func (self *CoreRunningTag) IsRunning() bool {
	return atomic.LoadInt64(&self.running) != 0
}

func (self *CoreRunningTag) SetRunning(v bool) {
	if v {
		atomic.StoreInt64(&self.running, 1)
	} else {
		atomic.StoreInt64(&self.running, 0)
	}
}

func (self *CoreRunningTag) WaitStopFinish() {
	self.stoppingWg.Wait()
}

func (self *CoreRunningTag) IsStopping() bool {
	return atomic.LoadInt64(&self.stopping) != 0
}

func (self *CoreRunningTag) StartStopping() {
	self.stoppingWg.Add(1)
	atomic.StoreInt64(&self.stopping, 1)
}

func (self *CoreRunningTag) EndStopping() {
	if self.IsStopping() {
		self.stoppingWg.Done()
		atomic.StoreInt64(&self.stopping, 0)
	}
}
