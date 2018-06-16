package component

import (
	"container/list"
	"github.com/zouyx/gopt/build"
	"github.com/zouyx/gopt/input"
)

var (
	buildGithubHandlers *list.List
)

func init() {
	buildGithubHandlers=list.New()
	buildGithubHandlers.PushBack((&build.DirBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.IgnoreBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.CoverageBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.BuildBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.PackageBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.ReadMeBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.TravisBuilder{}).Build)
	buildGithubHandlers.PushBack((&build.MainFileBuilder{}).Build)
}

func BuildGithubProject(params *input.Params) {
	processHandler(params,buildGithubHandlers)
}
