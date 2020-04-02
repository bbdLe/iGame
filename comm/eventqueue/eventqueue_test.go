package eventqueue

import (
	"fmt"
	"testing"
)

var (
	q *eventQueue
)

func TestEventQueue_Post(t *testing.T) {
	q.Post(func() {
		fmt.Println("Hello World")
	})
}

func TestEventQueue_StopLoop(t *testing.T) {
	q.StopLoop()
}

func TestEventQueue_Wait(t *testing.T) {
	q.Wait()
}

func TestMain(m *testing.M) {
	q = NewEventQueue()
	q.StartLoop()
	m.Run()
}
