package pipe

import (
	"sync"
)

type Pipe struct {
	mutex sync.Mutex
	cond *sync.Cond
	msgList []interface{}
}

func NewPipe() *Pipe {
	p := &Pipe{}
	p.cond = sync.NewCond(&p.mutex)

	return p
}

// 插入信息
func (p *Pipe) Add(msg interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.msgList = append(p.msgList, msg)
	p.cond.Signal()
}

func (p *Pipe) clear() {
	p.msgList = p.msgList[0:0]
}

// 阻塞获取信息
func (p *Pipe) Pick(msgList *[]interface{}) (exit bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// 没信息等信号
	if len(p.msgList) <= 0 {
		p.cond.Wait()
	}

	for i := range p.msgList {
		// 收到退出之后, 之后就不再收包了
		if p.msgList[i] == nil {
			exit = true
			break
		}

		*msgList = append(*msgList, p.msgList[i])
	}

	// 清空
	p.clear()

	return
}

func (p *Pipe) Reset() {
	p.msgList = p.msgList[0:0]
}