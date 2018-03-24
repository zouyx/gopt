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
	handlers.PushBack(getProjectType)
}

func GetParam() *Params{
	params:=&Params{
		ProjectType:1,
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
	1:github
	2:normal
	 */
	ProjectType int8 `json:"project_type"`
}