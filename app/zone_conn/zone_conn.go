package zone_conn

import (
	"github.com/bbdLe/iGame/app/zone_conn/logic/backend"
	"github.com/bbdLe/iGame/app/zone_conn/logic/frontend"
)

func Run() {
	backend.ConnectBackend()
	frontend.StartFrontEnd()

	select {

	}
}