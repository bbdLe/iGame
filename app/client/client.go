package client

import (
	"bufio"
	"fmt"
	"github.com/bbdLe/iGame/app/client/handler"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func connectConn(addr string, token string) error {
	var wg sync.WaitGroup
	wg.Add(1)

	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "client", addr, q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("session connect")
			ev.Session().Send(&proto.VerifyReq{
				Token: token,
				Server: "1",
			})
			wg.Done()
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
		default:
			handler.MsgDispatcher.OnEvent(ev)
		}
	})
	q.StartLoop()
	p.Start()

	wg.Wait()

	go func() {
		for {
			q.Post(func() {
				p.(peer.TCPConnector).Session().Send(&proto.HeartBeatReq{
				})
			})
			time.Sleep(time.Second)
		}
	}()

	readFromConsole(p)

	return nil
}

func readFromConsole(p comm.Peer) {
	reader := bufio.NewReader(os.Stdin)
	for {
		gm, err := reader.ReadString('\n')
		if err != nil {
			log.Logger.Debug("???")
			break
		}

		gm = strings.TrimRight(gm, "\n\r")

		gmList := strings.Split(gm, " ")
		if gmList[0] == "move" {
			x, err := strconv.ParseInt(gmList[1], 10, 64)
			if err != nil {
				log.Logger.Debug("???")
				continue
			}

			y, err := strconv.ParseInt(gmList[2], 10, 64)
			if err != nil {
				log.Logger.Debug(fmt.Sprintf("%v", err))
				continue
			}

			p.(interface{ Session() comm.Session}).Session().Send(&proto.MovePosReq{
				Pos:  &proto.Pos{
					X: x,
					Y: y,
				},
			})
		}
	}
}

func Run() {
	connectConn("localhost:10086", "token")
}