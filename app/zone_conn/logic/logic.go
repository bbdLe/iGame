package logic

const (
	AuthTimeKey = "auth_time"
	AuthKey = "auth"
	HeartBeatKey = "heart_beat"
	SessionTypeKey = "session_type"
)

var (
	BackEndMgr *BackEndManager
	FrontEndMgr *FrontEndManager
)

func init() {
	FrontEndMgr = NewFrontEventManager()
	BackEndMgr = NewBackEndManager()
}
