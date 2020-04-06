package processor

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
)

type ProcessorBinder func(bundle ProcessorBundle, cb EventCallback, args ...interface{})

var (
	procByName map[string]ProcessorBinder
)

func init() {
	procByName = make(map[string]ProcessorBinder)
}

func RegProcessor(name string, f ProcessorBinder) {
	if _, ok := procByName[name]; ok {
		panic(fmt.Errorf("Reg"))
	}

	procByName[name] = f
}

func GetProcessorList() []string {
	var names []string
	for n := range procByName {
		names = append(names, n)
	}

	return names
}

func BindProcessorHandler(peer comm.Peer, procName string, cb EventCallback, args ...interface{}) {
	if proc, ok := procByName[procName]; ok {
		bundle := peer.(ProcessorBundle)
		proc(bundle, cb, args)
	} else {
		panic(fmt.Errorf("BindProcessorHandler faild, %s not in handler", procName))
	}
}
