package logic

import (
	"github.com/bbdLe/iGame/comm"
)

const (
	AuthTimeKey = "auth_time"
	AuthKey = "auth"
	HeartBeatKey = "heart_beat"
	SessionTypeKey = "session_type"
)

var (
	BackEndMgr BackendManager
	FrontEndMgr FrontendManager
)

type BackendManager interface {
	Start()

	Send(interface{})
}

type FrontendManager interface {
	Start()

	GetSession(int64) (comm.Session, bool)
	AddSession(comm.Session)
	DelSession(comm.Session)

	Kick(int64)

	KickAll()
}