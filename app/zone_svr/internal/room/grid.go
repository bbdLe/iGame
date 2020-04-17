package room

import (
	"fmt"

	"github.com/bbdLe/iGame/app/zone_svr/internal"
)

type Grid struct {
	GID int64
	MinX int64
	MaxX int64
	MinY int64
	MaxY int64

	EntityMap map[internal.Entity]struct{}
}

func (self *Grid) AddEntity(e internal.Entity) {
	self.EntityMap[e] = struct{}{}
}

func (self *Grid) RemoveEntity(e internal.Entity) {
	delete(self.EntityMap, e)
}

func (self *Grid) GetEntities() []internal.Entity {
	var entities []internal.Entity
	for e := range self.EntityMap {
		entities = append(entities, e)
	}

	return entities
}

func (self *Grid) String() string {
	s := fmt.Sprintf("Grid(%d)\n MinX: %d, MaxX: %d\n MinY: %d, MaxY: %d\n", self.GID,
		self.MinX, self.MaxX, self.MinY, self.MaxY)
	return s
}

func NewGrid(gid, minx, maxx, miny, maxy int64) *Grid {
	return &Grid{
		GID: gid,
		MinX: minx,
		MaxX: maxx,
		MinY: miny,
		MaxY: maxy,
		EntityMap: make(map[internal.Entity]struct{}),
	}
}