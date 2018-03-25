package build

import (
	"github.com/zouyx/gopt/input"
	"os"
	"fmt"
	"github.com/zouyx/gopt/message"
)

const (
	FULL_PATH ="%v/src/main"
)

// based input params get project full path
func getFullPath(params *input.Params) string {
	return fmt.Sprintf(FULL_PATH, params.ProjectName)
}

type DirBuilder struct {

}

// dir build method
func (this *DirBuilder) Build(params *input.Params) {
	fullPath := getFullPath(params)
	err := os.MkdirAll(fullPath, os.ModePerm)

	if err!=nil{
		message.FormatError(func() {
			panic(fmt.Sprintf("Create dir fail! error:%v",err.Error()))
		})
	}
	message.Success(fmt.Sprintf("created %v",fullPath))
}
