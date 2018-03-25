package build

import (
	"github.com/zouyx/gopt/input"
	"os"
	"fmt"
	"github.com/zouyx/gopt/message"
)

const (
	FULL_PATH ="%v/src/%v"
)

type DirBuilder struct {

}

// dir build method
func (this *DirBuilder) Build(params *input.Params) {
	err := os.MkdirAll(fmt.Sprintf(FULL_PATH,params.ProjectName,params.ProjectName), os.ModePerm)

	if err!=nil{
		message.FormatError(func() {
			panic(fmt.Sprintf("Create dir fail! error:%v",err.Error()))
		})
	}
}
