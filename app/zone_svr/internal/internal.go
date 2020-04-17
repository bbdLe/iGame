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

type Player interface {
	Init()

	Tick()

	Session() comm.Session

	ID() int64
	Name() string

	OnLogout()
	OnLogin()

	SetHeartBeat(time.Time)
	HeartBeat() time.Time

	Room() Room
	SetRoom(Room)

	BaseInfo() PlayerBaseInfo

	Send(interface{})
}


type GameManager interface{
	Start()

	KickPlayer(Player)
	SetPlayer(int64, Player)
	GetPlayer(int64) (Player, bool)
	DelPlayer(int64)

	VisitPlayer(func(Player))

	OnPlayerLogin(Player)
}

type Room interface {
	ID() int64

	SetID(int64)

	AddPlayer(Player)

	RemovePlayer(Player)

	Broadcast(interface{})
}

type RoomManager interface {
	Tick()

	NewRoom() Room

	GetRoom(int64) (Room, bool)
}

type PlayerBaseInfo interface {
	Exp() int64

	Level() int

	AddExp(int64)

	Player() Player
}

func Send2Player(player Player, msg interface{}) {
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
