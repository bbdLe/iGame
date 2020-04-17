package room

import (
	"reflect"
	"testing"
)

var (
	aoiMgr *AoiManager
)

func TestAoiManager_GridLength(t *testing.T) {
	if aoiMgr.GridLength() != 50 {
		t.Errorf("aoiMgr length is %d, but actual is %d", aoiMgr.GridLength(), 50)
	}
}

func TestAoiManager_GridWidth(t *testing.T) {
	if aoiMgr.GridWidth() != 50 {
		t.Errorf("aoiMgr length is %d, but actual is %d", aoiMgr.GridLength(), 50)
	}
}

func TestAoiManager_GetSurroundGridsByGid(t *testing.T) {
	grids := aoiMgr.GetSurroundGridsByGid(0)


	except := []int64{0, 1, 5, 6}
	var res []int64
	for _, g := range grids {
		//log.Println(g)
		res = append(res, g.GID)
	}
	if !reflect.DeepEqual(except, res) {
		t.Errorf("except is %v, actual is %v", except, res)
	}
}

func TestMain(m *testing.M) {
	aoiMgr = NewAoiManager(0, 250, 5, 0, 250, 5)
	m.Run()
}
