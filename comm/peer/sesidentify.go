package peer

type CoreSessionIdentify struct {
	id int64
}

func (self *CoreSessionIdentify) SetID(pId int64) {
	self.id = pId
}

func (self *CoreSessionIdentify) ID() int64 {
	return self.id
}
