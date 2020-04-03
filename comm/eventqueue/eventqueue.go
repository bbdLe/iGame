package eventqueue

import (
	"github.com/bbdLe/iGame/comm/pipe"
	"github.com/bbdLe/iGame/comm/session"
	"log"
	"sync"
)

type EventQueue interface {
	StartLoop()

	StopLoop()

	Wait()

	Post(func ())
}

type PaincCbFunc func(interface{}, EventQueue)

type eventQueue struct {
	*pipe.Pipe
	wg sync.WaitGroup
	paincCb PaincCbFunc
}

func defaultPaincCallback(err interface{}, q EventQueue) {
	log.Println("eventqueue error, ", err)
}

func NewEventQueue() *eventQueue {
	return &eventQueue{
		Pipe: pipe.NewPipe(),
		paincCb :defaultPaincCallback,
	}
}

func (q *eventQueue) StartLoop() {
	q.wg.Add(1)

	go func() {
		var cbList []interface{}

		for {
			cbList = cbList[0:0]

			exit := q.Pick(&cbList)
			for _, cb := range cbList {
				switch t := cb.(type) {
				case func():
					q.safeCall(t)
				case nil:
					break
				default :
					log.Printf("wrong type %T", t)
				}
			}

			if exit {
				break
			}
		}

		q.wg.Done()
	}()
}

func (q *eventQueue) StopLoop() {
	q.Add(nil)
}

func (q *eventQueue) Wait() {
	q.wg.Wait()
}

func (q *eventQueue) Post(cb func()) {
	q.Add(cb)
}

func (q *eventQueue) safeCall(cb func()) {
	defer func() {
		if err := recover(); err != nil {
			q.paincCb(err, q)
		}
	}()

	cb()
}

func SessionCall(sess session.Session, cb func()) {
	q := sess.Peer().(interface {
		Queue() EventQueue
	}).Queue()

	QueueCall(q, cb)
}

func QueueCall(q EventQueue, cb func()) {
	if q == nil {
		cb()
	} else {
		q.Post(cb)
	}
}