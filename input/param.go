package input

import (
	"container/list"
)

var (
	handlers *list.List
)

func init() {
	handlers=list.New()
	handlers.PushBack(welcome)
	handlers.PushBack(initProjectType)
}

// get all params for build project structure
func GetAllParams() *Params{
	i, _ := GetDefaultProjectType()
	params:=&Params{
		ProjectType:i,
	}
	for e := handlers.Front(); e != nil; e = e.Next() {
		e.Value.(func(*Params))(params)
	}
	return params
}

type Params struct {
	/**
	project type
	0:custom
	1:normal
	2:github
	 */
	ProjectType int8 `json:"project_type"`
}