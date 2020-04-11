package websocket

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"net"
	"net/http"

	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/gorilla/websocket"
)

type wsAcceptor struct {
	peer.CoreSessionManager
	peer.CorePeerProperty
	peer.CoreContextSet
	peer.CoreProcBundle
	peer.CoreRunningTag

	certfile string
	keyfile string

	upgrader websocket.Upgrader
	listener net.Listener
	sv *http.Server
}

func (self *wsAcceptor) SetUpgrader(upgrader interface{}) {
	self.upgrader = upgrader.(websocket.Upgrader)
}

func (self *wsAcceptor) SetHttps(certfile, keyfile string) {
	self.certfile = certfile
	self.keyfile = keyfile
}

func (self *wsAcceptor) Port() int {
	if self.listener == nil {
		return 0
	} else {
		return self.listener.Addr().(*net.TCPAddr).Port
	}
}

func (self *wsAcceptor) IsReady() bool {
	return self.Port() != 0
}

func (self *wsAcceptor) Start() {
	self.WaitStopFinish()

	if self.IsRunning() {
		return
	}

	var err error
	self.listener, err = net.Listen("tcp", self.Address())
	if err != nil {
		log.Logger.Error(fmt.Sprintf("listen fail : %v", err))
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := self.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Logger.Error(fmt.Sprintf("upgrade fail : %v", err))
			return
		}

		ses := newSession(c, self, nil)
		ses.SetContext("request", r)
		ses.Start()
		self.ProcessEvent(&event.RecvMsgEvent{Ses: ses, Msg : &sysmsg.SessionAccepted{}})
	})

	self.sv = &http.Server{Addr: self.Address(), Handler : mux}

	go func() {
		log.Logger.Info("Start Listen")
		self.SetRunning(true)
		err := self.sv.Serve(self.listener)
		if err != nil {
			log.Logger.Error(fmt.Sprintf("Listen fail : %v", err))
		}
		self.SetRunning(false)
		self.EndStopping()
	}()
}

func (self *wsAcceptor) Stop() {
	if !self.IsRunning() {
		return
	}

	if self.IsStopping() {
		return
	}

	self.StartStopping()
	self.listener.Close()
	self.CloseAllSession()
	self.WaitStopFinish()
}

func (self *wsAcceptor) TypeName() string {
	return "ws.Acceptor"
}

func init() {
	peer.RegPeerCreateor(func() comm.Peer{
		p := &wsAcceptor{
			upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		}
		return p
	})
}