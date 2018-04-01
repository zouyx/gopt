package build

import (
	"github.com/zouyx/gopt/input"
	"fmt"
	"os"
	"github.com/zouyx/gopt/message"
	"bufio"
)

type StructureBuilder interface {
	Build(param *input.Params)
}


const (
	SRC_PATH ="%v/src"
	FULL_PATH =SRC_PATH+"/main"
)

// based input params get project full path
func getFullPath(params *input.Params) string {
	return fmt.Sprintf(FULL_PATH, params.ProjectName)
}

// based input params get project full path
func getSrc(params *input.Params) string {
	return fmt.Sprintf(SRC_PATH, params.ProjectName)
}

// write to file name
func write(fileName,content string)  {
	f, err := os.Create(fileName)
	if err!=nil{
		message.FormatError(func() {
			panic(fmt.Sprintf("Create %v fail! error:%v",fileName,err.Error()))
		})
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	if err!=nil{
		message.FormatError(func() {
			panic(fmt.Sprintf("Write %v fail! error:%v",fileName,err.Error()))
		})
	}
	w.Flush()
	f.Close()

	message.Success(fmt.Sprintf("Created %v",fileName))
}