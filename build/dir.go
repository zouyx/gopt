package build

import (
	"github.com/zouyx/gopt/input"
	"os"
	"fmt"
	"github.com/zouyx/gopt/message"
)

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
