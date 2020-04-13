package player

type PlayerCmpt interface {
	Tick()

	Init()
}

var (
	PlayerMap map[int64]*Player
)

type Player struct {
	baseInfo PlayerBaseInfo

	Cmpts []PlayerCmpt
	SessionID int64
	Status int32
}

func (self *Player) RegCmpt(m PlayerCmpt) {
	self.Cmpts = append(self.Cmpts, m)
}

func (self *Player) Init() {
	self.RegCmpt(&self.baseInfo)
}

func (self *Player) Tick() {
	for _, cmpt := range self.Cmpts {
		cmpt.Tick()
	}
}

func NewPlayer(sessionID int64) *Player {
	self := &Player{
		SessionID: sessionID,
	}
	self.Init()
	return self
}

func SetPlayer(sessionID int64, player *Player) {
	PlayerMap[sessionID] = player
}

func GetPlayer(sessionID int64) (*Player, bool) {
	p, ok := PlayerMap[sessionID]
	return p, ok
}

func init() {
	PlayerMap = make(map[int64]*Player)
}
