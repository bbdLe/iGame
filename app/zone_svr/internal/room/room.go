package room

import "github.com/bbdLe/iGame/app/zone_svr/internal"

type Room struct {
	id int64

	playerList []internal.CommPlayer
}

func (self *Room) ID() int64 {
	return self.id
}

func (self *Room) SetID(id int64) {
	self.id = id
}

func (self *Room) Tick() {

}

func (self *Room) VisitPlayer(f func(p internal.CommPlayer)) {
	for _, p := range self.playerList {
		f(p)
	}
}