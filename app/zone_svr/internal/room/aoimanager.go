package room

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_svr/internal"
)

type AoiManager struct {
	MinX int64
	MaxX int64
	MinY int64
	MaxY int64

	XNums int64
	YNums int64

	grids map[int64]*Grid
}

func (self *AoiManager) GridWidth() int64 {
	return (self.MaxX - self.MinX) / self.XNums
}

func (self *AoiManager) GridLength() int64 {
	return (self.MaxY - self.MinY) / self.YNums
}

func (self *AoiManager) String() string {
	s := fmt.Sprintf("AOIManagr:\nminX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOI Manager:\n",
		self.MinX, self.MaxX, self.XNums, self.MinY, self.MaxY, self.YNums)
	for _,grid := range self.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

// 寻找九宫格
func (self *AoiManager) GetSurroundGridsByGid(gid int64) []*Grid {
	// 先检查当前网格
	curGrid, ok := self.grids[gid]
	if !ok {
		return nil
	}

	var grids []*Grid
	// 当前网格先加入
	grids = append(grids, curGrid)

	// x轴index
	idx := gid % self.XNums

	// 检查左边
	if idx > 0 {
		grids = append(grids, self.grids[gid - 1])
	}
	// 检查右边
	if idx < self.XNums - 1 {
		grids = append(grids, self.grids[gid + 1])
	}

	// y轴上下都检查一次
	for _, g := range grids {
		id := g.GID

		idy := id / self.XNums

		// 上面有一层
		if idy > 0 {
			grids = append(grids, self.grids[id - self.XNums])
		}

		// 下面有一层
		if idy < self.YNums - 1 {
			grids = append(grids, self.grids[id + self.XNums])
		}
	}

	return grids
}

func (self *AoiManager) GetGridIDByPos(x, y int64) int64 {
	idx := (x - self.MinX) / self.GridWidth()
	idy := (y - self.MinY) / self.GridLength()

	return idy * self.XNums + idx
}

func (self *AoiManager) GetEntryByPos(x, y int64) []internal.Entity {
	gid := self.GetGridIDByPos(x, y)
	_, ok := self.GetGridByGID(gid)
	if !ok {
		return nil
	}

	grids := self.GetSurroundGridsByGid(gid)
	var res []internal.Entity
	for _, g := range grids {
		entities := g.GetEntities()
		for _, e := range entities {
			res = append(res, e)
		}
	}

	return res
}

func (self *AoiManager) GetEntryByGID(gid int64) []internal.Entity {
	grids := self.GetSurroundGridsByGid(gid)
	var res []internal.Entity
	for _, g := range grids {
		entities := g.GetEntities()
		for _, e := range entities {
			res = append(res, e)
		}
	}

	return res
}

func (self *AoiManager) AddEntityByGID(entity internal.Entity, gid int64) {
	grid, ok := self.GetGridByGID(gid)
	if !ok {
		return
	}

	grid.AddEntity(entity)
}

func (self *AoiManager) AddEntityByPos(entity internal.Entity, x, y int64) {
	gid := self.GetGridIDByPos(x, y)
	grid, ok := self.GetGridByGID(gid)
	if !ok {
		return
	}

	grid.AddEntity(entity)
}

func (self *AoiManager) RemoveEntityByGID(entity internal.Entity, gid int64) {
	grid, ok := self.GetGridByGID(gid)
	if !ok {
		return
	}

	grid.RemoveEntity(entity)
}

func (self *AoiManager) RemoveEntityByPos(entity internal.Entity, x, y int64) {
	gid := self.GetGridIDByPos(x, y)
	grid, ok := self.GetGridByGID(gid)
	if !ok {
		return
	}

	grid.RemoveEntity(entity)
}

func (self *AoiManager) GetGridByGID(gid int64) (*Grid, bool) {
	grid, ok := self.grids[gid]
	return grid, ok
}

func NewAoiManager(minx, maxx, xnums, miny, maxy, ynums int64) *AoiManager {
	aoiMgr := &AoiManager{
		MinX: minx,
		MaxX: maxx,
		MinY: minx,
		MaxY: maxy,
		XNums: xnums,
		YNums: ynums,
		grids: make(map[int64]*Grid),
	}

	for i := int64(0); i < xnums; i++ {
		for j := int64(0); j < ynums; j++ {
			gridId := j * xnums + i

			aoiMgr.grids[gridId] = NewGrid(gridId,
				aoiMgr.MinX + i * aoiMgr.GridWidth(),
				aoiMgr.MinX + (i + 1) * aoiMgr.GridWidth(),
				aoiMgr.MinY + j * aoiMgr.GridLength(),
				aoiMgr.MinY + (j + 1) * aoiMgr.GridLength())
		}
	}

	return aoiMgr
}
