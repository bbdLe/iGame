package pipe

import (
	"fmt"
	"testing"
)

var (
	p *Pipe
)

func TestPipe_Add(t *testing.T) {
}

func TestPipe_Pick(t *testing.T) {
	var msgList []interface{}
	exit := p.Pick(&msgList)
	if len(msgList) != 2 {
		t.Errorf("len(msgList) != except")
	}

	for _, msg := range msgList {
		msg.(func())()
	}

	if exit != true {
		t.Errorf("can't get exit sign")
	}
}

func TestMain(m *testing.M) {
	p = NewPipe()
	p.Add(func() {
		fmt.Println("Hello World")
	})
	p.Add(func() {
		fmt.Println("Good Bye")
	})
	p.Add(nil)
	m.Run()
}