package zone_conn

import (
	"github.com/bbdLe/iGame/app/zone_conn/logic"
)

func Run() {
	logic.BackEndMgr.Start()
	logic.FrontEndMgr.Start()

	select {

	}
}