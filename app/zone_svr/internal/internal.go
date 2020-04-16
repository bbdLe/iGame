package internal

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/proto"
	"time"
)

var (
	ZoneAcceptor comm.Peer
	ZoneEventQueue comm.EventQueue
	GameMgr GameManager
	RoomMgr RoomManager
)

type CommPlayer interface {
	Init()

	Tick()

	Session() comm.Session

	ID() int64
	Name() string

	OnLogout()
	OnLogin()

	SetHeartBeat(time.Time)
	HeartBeat() time.Time

	Room() CommRoom
	SetRoom(CommRoom)

	BaseInfo() CommPlayerBaseInfo

	Send(interface{})
}


type GameManager interface{
	Start()

	KickPlayer(CommPlayer)

	SetPlayer(int64, CommPlayer)

	GetPlayer(int64) (CommPlayer, bool)

	DelPlayer(int64)

	VisitPlayer(func(CommPlayer))
}

type CommRoom interface {
	ID() int64

	SetID(int64)

	AddPlayer(CommPlayer)

	RemovePlayer(CommPlayer)

	Broadcast(interface{})
}

type RoomManager interface {
	Tick()

	NewRoom() CommRoom

	GetRoom(int64) (CommRoom, bool)
}

type CommPlayerBaseInfo interface {
	Exp() int64

	Level() int

	AddExp(int64)

	Player() CommPlayer
}

func Send2Player(player CommPlayer, msg interface{}) {
	meta := comm.MessageMetaByMsg(msg)
	if meta == nil {
		log.Logger.Error("get meta fail")
		return
	}

	data, err := meta.Codec.Encode(msg, nil)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Encode msg(%d) fail: %v", meta.MsgId, err))
		return
	}

	player.Session().Send(&proto.TransmitRes{
		MsgId: int32(meta.MsgId),
		MsgData: data.([]byte),
		ClientId: player.ID(),
	})
}
