package component

import (
	"github.com/zouyx/gopt/input"
	"container/list"
)


func processHandler(params *input.Params,buildHandlers *list.List)  {
	for e := buildHandlers.Front(); e != nil; e = e.Next() {
		e.Value.(func(*input.Params))(params)
	}
}