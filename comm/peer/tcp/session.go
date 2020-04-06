package tcp

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/pipe"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/comm/util"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type tcpSession struct {
	peer.CoreContextSet
	peer.CoreSessionIdentify
	*peer.CoreProcBundle

	pInterface comm.Peer

	conn net.Conn
	connGuard sync.RWMutex

	exitSync sync.WaitGroup

	// 发送队列
	sendQueue *pipe.Pipe

	cleanupGuard sync.Mutex

	endNotify func()

	closing int64
}

func (self *tcpSession) SetConn(conn net.Conn) {
	self.connGuard.Lock()
	defer self.connGuard.Unlock()

	self.conn = conn
}

func (self *tcpSession) Conn() net.Conn {
	self.connGuard.RLock()
	defer self.connGuard.RUnlock()

	return self.conn
}

func (self *tcpSession) Peer() comm.Peer {
	return self.pInterface
}

func (self *tcpSession) Raw() interface{} {
	return self.Conn()
}

func (self *tcpSession) Close() {
	bClose := atomic.CompareAndSwapInt64(&self.closing, 0, 1)
	if !bClose {
		return
	}

	conn := self.Conn()
	if conn != nil {
		tcpConn := conn.(*net.TCPConn)
		tcpConn.CloseRead()
		tcpConn.SetReadDeadline(time.Now())
	}
}

func (self *tcpSession) Send(msg interface{}) {
	if msg == nil {
		return
	}

	// 关闭之后不再发送
	if self.IsClose() {
		return
	}

	self.sendQueue.Add(msg)
}

func (self *tcpSession) IsClose() bool {
	return atomic.LoadInt64(&self.closing) != 0
}

func (self *tcpSession) ProtectedReadMessage() (msg interface{}, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("IO panic : %s", err)
			self.Close()
		}
	}()

	return self.ReadMessage(self)
}

func (self *tcpSession) recvLoop() {
	var capturePanic bool

	if i, ok := self.Peer().(comm.CaptureIOPanic); ok {
		capturePanic = i.CaptureIOPanic()
	}

	for self.Conn() != nil {
		var msg interface{}
		var err error

		if capturePanic {
			msg, err = self.ProtectedReadMessage()
		} else {
			msg, err = self.ReadMessage(self)
		}

		if err != nil {
			if !util.IsEOFOrNetReadError(err) {
				log.Println("session closed, sesid : %d, err := %s", self.ID(), err)
			}

			self.sendQueue.Add(nil)

			closedMsg := &sysmsg.SessionClose{}
			if self.IsClose() {
				closedMsg.Reason = sysmsg.CloseReason_IO
			}

			self.ProcessEvent(&event.RecvMsgEvent{Ses : self, Msg : closedMsg})
			break
		}

		self.ProcessEvent(&event.RecvMsgEvent{Ses : self, Msg: msg})
	}

	self.exitSync.Done()
}

func (self *tcpSession) sendLoop() {
	var writeList []interface{}

	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList)

		for _, msg := range writeList {
			self.SendMessage(&event.SendMsgEvent{Msg : msg, Ses : self})
		}

		if exit {
			break
		}
	}

	conn := self.Conn()
	if conn != nil {
		conn.Close()
	}

	self.exitSync.Done()
}

func (self *tcpSession) Start() {
	atomic.StoreInt64(&self.closing, 0)
	self.sendQueue.Reset()
	self.exitSync.Add(2)

	self.Peer().(peer.SessionManager).Add(self)

	// 清理
	go func() {
		self.exitSync.Wait()

		self.Peer().(peer.SessionManager).Remove(self)
		if self.endNotify != nil {
			self.endNotify()
		}
	}()

	// 接收协程
	self.recvLoop()

	// 发送协程
	self.sendLoop()
}

func newSession(conn net.Conn, p comm.Peer, endNotify func()) *tcpSession {
	self := &tcpSession{
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