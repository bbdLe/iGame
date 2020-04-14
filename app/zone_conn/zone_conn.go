package zone_conn

import (
	"github.com/bbdLe/iGame/app/zone_conn/handler"
	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func Run() {
	handler.ConnectBackend()
	handler.StartFroneEnd()

	select {

	}
}