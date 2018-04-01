package component

import (
	"container/list"
	"github.com/zouyx/gopt/build"
	"github.com/zouyx/gopt/input"
)

var (
	buildHandlers *list.List
)

func init() {
	buildHandlers=list.New()
	buildHandlers.PushBack((&build.DirBuilder{}).Build)
	buildHandlers.PushBack((&build.MainFileBuilder{}).Build)
	buildHandlers.PushBack((&build.IgnoreBuilder{}).Build)
}

func BuildNormalProject(params *input.Params) {
	processHandler(params,buildHandlers)
}
