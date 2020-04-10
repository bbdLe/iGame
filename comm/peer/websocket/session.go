package websocket

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"sync"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/pipe"
	"github.com/bbdLe/iGame/comm/util"

	"github.com/gorilla/websocket"
)

type wsSession struct {
	peer.CoreContextSet
	peer.CoreSessionIdentify
	*peer.CoreProcBundle

	pInterface comm.Peer
	conn *websocket.Conn

	exitSync sync.WaitGroup
	sendQueue *pipe.Pipe
	endNotify func()
}

func (self *wsSession) Peer() comm.Peer {
	return self.pInterface
}

func (self *wsSession) Raw() interface{} {
	return self.conn
}

func (self *wsSession) Close() {
	self.sendQueue.Add(nil)
}

func (self *wsSession) Send(msg interface{}) {
	self.sendQueue.Add(msg)
}

func (self *wsSession) Start() {
	self.Peer().(peer.SessionManager).Add(self)
	self.exitSync.Add(2)

	go func() {
		self.exitSync.Wait()

		self.Peer().(peer.SessionManager).Remove(self)
		if self.endNotify != nil {
			self.endNotify()
		}
	}()

	go self.recvLoop()
	go self.sendLoop()
}

func (self *wsSession) recvLoop() {
	defer self.exitSync.Done()

	for {
		msg, err := self.ReadMessage(self)
		if err != nil {
			if !util.IsEOFOrNetReadError(err) {
				log.Logger.Error(fmt.Sprintf("Recv fail : %v", err))
			}

			self.ProcessEvent(&event.RecvMsgEvent{self, &sysmsg.SessionClose{}})
			break
		}

		self.ProcessEvent(&event.RecvMsgEvent{Ses : self, Msg : msg})
	}

	self.Close()
}

func (self *wsSession) sendLoop() {
	defer self.exitSync.Done()

	var writeList []interface{}
	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList)

		for _, msg := range writeList {
			self.ProcessEvent(&event.SendMsgEvent{Ses : self, Msg : msg})
		}

		if exit {
			break
		}
	}

	if self.conn != nil {
		self.conn.Close()
		self.conn = nil
	}
}

func newSession(conn *websocket.Conn, p comm.Peer, endNotify func()) *wsSession {
	self := &wsSession{
		conn: conn,
		endNotify: endNotify,
		sendQueue: pipe.NewPipe(),
		pInterface: p,
		CoreProcBundle: p.(interface{
			GetBundle() *peer.CoreProcBundle
		}).GetBundle(),
	}

	return self
}