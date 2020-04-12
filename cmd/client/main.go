package main

import (
	"fmt"
	"sync"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func GetConnAddr() (string, string, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	connAddr := ""
	token := ""

	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "client", "localhost:12315", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("session connect")
			// send login
			req := &proto.LoginReq{
				Version: "0.0.1",
				Platform : "HuaWei",
				Uid : "sdadasdada",
			}
			ev.Session().Send(req)
		case *proto.LoginRes:
			connAddr = msg.GetServerAddr()
			token = msg.GetToken()
			ev.Session().Close()
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
			wg.Done()
		}
	})
	q.StartLoop()
	p.Start()

	wg.Wait()
	return connAddr, token, nil
}

func connectConn(addr string, token string) error {
	var wg sync.WaitGroup
	wg.Add(1)

	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "client", addr, q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("session connect")
			ev.Session().Send(&proto.VerifyReq{
				Token: token,
				Server: "1",
			})
		case *proto.VerifyRes:
			log.Logger.Debug(fmt.Sprintf("Recv VerifyRes : %v", msg.GetRetCode()))
			ev.Session().Close()
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
			wg.Done()
		}
	})
	q.StartLoop()
	p.Start()

	wg.Wait()

	return nil
}

func main() {
	addr, token, err := GetConnAddr()
	if err != nil {
		log.Logger.Error(fmt.Sprintf("GetConnAddr Fail: %s", err))
		return
	}
	log.Logger.Debug(addr)
	log.Logger.Debug(token)
	if addr == "" || token == "" {
		return
	}
	err = connectConn(addr, token)
	if err != nil {
		log.Logger.Error("connectConn fail")
	}
}
