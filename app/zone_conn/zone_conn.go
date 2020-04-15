package zone_conn

import (
	"github.com/bbdLe/iGame/app/zone_conn/handler"
)

func Run() {
	handler.ConnectBackend()
	handler.StartFroneEnd()

	select {

	}
}