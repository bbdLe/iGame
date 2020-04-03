package peer

type SessionIdentify struct {
	id int64
}

func (self *SessionIdentify) SetID(pId int64) {
	self.id = pId
}

func (self *SessionIdentify) GetID() int64 {
	return self.id
}
