package room

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/proto"
	"time"
)

type RoomImpl struct {
	aoiMgr *AoiManager
	Id int64

	playerMap map[int64]internal.Player

	canFree bool
	emptyTime time.Time
}

func (self *RoomImpl) ID() int64 {
	return self.Id
}

func (self *RoomImpl) SetID(id int64) {
	self.Id = id
}

func (self *RoomImpl) Tick() {
}

func (self *RoomImpl) AddPlayer(p internal.Player) {
	log.Logger.Info(fmt.Sprintf("player[%d] enter room[%d]", p.ID(), self.ID()))

	self.playerMap[p.ID()] = p
	p.SetRoom(self)
	self.OnPlayerEnter(p)

	xPos := int64(0)
	yPos := int64(0)
	p.BaseInfo().Pos().Set(xPos, yPos)

	// 加入网格中
	self.aoiMgr.AddEntityByPos(p.(internal.Entity), xPos, yPos)

	// 向周围广播自己进入视野里面
	entities := self.aoiMgr.GetEntryByPos(xPos, yPos)
	for _, entity := range entities {
		entity.EnterView(p.(internal.Entity))

		// 自己看到自己一次就可以
		if entity.ID() != p.ID() {
			p.EnterView(entity)
		}
	}
}

func (self *RoomImpl) RemovePlayer(p internal.Player) {
	log.Logger.Info(fmt.Sprintf("player[%d] leave room[%d]", p.ID(), self.ID()))

	entities := self.aoiMgr.GetEntryByPos(p.BaseInfo().Pos().X(), p.BaseInfo().Pos().Y())
	for _, entity := range entities {
		entity.LeaveView(p.(internal.Entity))

		// 自己离开视野一次就可以
		if entity.ID() != p.ID() {
			p.LeaveView(entity)
		}
	}
	// 退出网格
	self.aoiMgr.RemoveEntityByPos(p.(internal.Entity), p.BaseInfo().Pos().X(), p.BaseInfo().Pos().Y())

	p.SetRoom(nil)
	delete(self.playerMap, p.ID())
	self.OnPlayerLeave(p)

	if self.Count() == 0 {
		self.emptyTime = time.Now()
	}
}

func (self *RoomImpl) Count() int {
	return len(self.playerMap)
}

func (self *RoomImpl) Empty() bool {
	return self.Count() == 0
}

func (self *RoomImpl) CanFree() bool {
	if self.Empty() && self.emptyTime.Add(time.Minute).Before(time.Now()) && self.canFree {
		return true
	} else {
		return false
	}
}

func (self *RoomImpl) OnPlayerMove(p internal.Player, x, y int64) {
	oldX := p.BaseInfo().Pos().X()
	oldY := p.BaseInfo().Pos().Y()
	p.BaseInfo().Pos().Set(x, y)

	// 获取老的网格
	oldGID := self.aoiMgr.GetGridIDByPos(oldX, oldY)
	newGID := self.aoiMgr.GetGridIDByPos(x, y)
	log.Logger.Debug(fmt.Sprintf("NewGID : %d, OldGID : %d", newGID, oldGID))

	// 网格更换
	if oldGID != newGID {
		// 老九宫格的人
		oldEntities := self.aoiMgr.GetEntryByGID(oldGID)
		// 新九宫格的人
		newEntities := self.aoiMgr.GetEntryByGID(newGID)
		// 需要移除视野的人
		var interseOldEntities []internal.Entity

		// 求出老九宫格出现, 但没在新九宫格的人, 先不考虑性能
		for _, e := range oldEntities {
			found := false

			// 自己不算进去
			if e.ID() == p.ID() {
				continue
			}

			for _, ne := range newEntities {
				if ne.ID() == e.ID() {
					found = true
				}
			}

			if !found {
				interseOldEntities = append(interseOldEntities, e)
			}
		}

		// 通知离开视野
		for _, e := range interseOldEntities {
			log.Logger.Debug(fmt.Sprintf("%d", e.ID()))
			e.LeaveView(p.(internal.Entity))
			p.LeaveView(e)
		}

		// 求出新视野中出现, 但不在老视野的人, 先不考虑性能
		var interseNewEntites []internal.Entity
		for _, ne := range newEntities {
			found := false
			for _, e := range oldEntities {
				if e.ID() == ne.ID() {
					found = true
				}
			}

			if !found {
				interseNewEntites = append(interseNewEntites, ne)
			}
		}

		// 通知进入视野
		for _, e := range interseNewEntites {
			e.EnterView(p.(internal.Entity))
			p.EnterView(e.(internal.Entity))
		}

		// 同步位置
		for _, e := range newEntities {
			e.SendPos(p.(internal.Entity))
		}

		// 把自己移除
		self.aoiMgr.RemoveEntityByGID(p.(internal.Entity), oldGID)
		// 加入新网格中
		self.aoiMgr.AddEntityByGID(p.(internal.Entity), newGID)
		// 发送自己位置
		p.SendPos(p.(internal.Entity))
	} else {
		entities := self.aoiMgr.GetEntryByGID(oldGID)
		// 网格不变的情况下, 只需要广播位置信息
		for _, entity := range entities {
			entity.SendPos(p.(internal.Entity))
		}
	}
}

func (self *RoomImpl) VisitPlayer(f func(p internal.Player)) {
	for _, p := range self.playerMap {
		f(p)
	}
}

func (self *RoomImpl) Broadcast(msg interface{}) {
	self.VisitPlayer(func(p internal.Player) {
		p.Send(msg)
	})
}

func (self *RoomImpl) OnPlayerEnter(p internal.Player) {
	var msg proto.BroadcastMsgRes
	msg.Msg = fmt.Sprintf("欢迎%s进入房间", p.Name())
	msg.Type = proto.MSG_TYPE_SYSTEM
	self.Broadcast(&msg)
}

func (self *RoomImpl) OnPlayerLeave(p internal.Player) {
	var msg proto.BroadcastMsgRes
	msg.Msg = fmt.Sprintf("玩家%s离开了房间", p.Name())
	msg.Type = proto.MSG_TYPE_SYSTEM
	self.Broadcast(&msg)
}

func (self *RoomImpl) SetCanFree(b bool) {
	self.canFree = b
}

func NewRoom(id int64) *RoomImpl {
	return &RoomImpl{
		Id : id,
		playerMap : make(map[int64]internal.Player),
		aoiMgr: NewAoiManager(0, 250, 5, 0, 250, 5),
		canFree: false,
	}
}