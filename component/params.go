package component


import (
	"container/list"
	"github.com/zouyx/gopt/input"
)

var (
	handlers *list.List
)

func init() {
	handlers=list.New()
	handlers.PushBack((&input.WelcomeHandler{}).Handle)
	handlers.PushBack((&input.ProjectNameInputHandler{}).Handle)
	handlers.PushBack((&input.ProjectTypeInputHandler{}).Handle)
}

// get all params for build project structure
func GetAllParams() *input.Params{
	i, _ := input.GetDefaultProjectType()
	params:=&input.Params{
		ProjectType:i,
	}
	for e := handlers.Front(); e != nil; e = e.Next() {
		e.Value.(func(*input.Params))(params)
	}
	return params
}
