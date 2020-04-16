package zone_conn

import (
	"github.com/bbdLe/iGame/app/zone_conn/logic"

	_"github.com/bbdLe/iGame/app/zone_conn/logic/backend"
	_"github.com/bbdLe/iGame/app/zone_conn/logic/frontend"
)

func Run() {
	logic.BackEndMgr.Start()
	logic.FrontEndMgr.Start()

	select {

	}
}