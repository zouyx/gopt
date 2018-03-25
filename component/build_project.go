package component

import (
	"github.com/zouyx/gopt/input"
	"container/list"
	"github.com/zouyx/gopt/build"
)

var (
	buildHandlers *list.List
)

func init() {
	buildHandlers=list.New()
	buildHandlers.PushBack((&build.DirBuilder{}).Build)
}

func BuildProject(params *input.Params) {
	for e := buildHandlers.Front(); e != nil; e = e.Next() {
		e.Value.(func(*input.Params))(params)
	}
}