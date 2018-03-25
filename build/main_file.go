package build

import (
	"github.com/zouyx/gopt/input"
	"os"
	"github.com/zouyx/gopt/message"
	"fmt"
	"bufio"
)

const (
	MAIN_FILE_TEMPLATE="package main\n" +
		"\n" +
		"import (\n" +
		"\t\"fmt\"\n" +
		")\n" +
		"\n" +
		"func main() {\n" +
		"\tfmt.Println(\"welcome to your awesome project: %v!\")\n" +
		"}\n"
)


type MainFileBuilder struct {

}

// dir build method
func (this *MainFileBuilder) Build(params *input.Params) {
	fullPath := getFullPath(params)
	mainFile := fullPath + "/main.go"
	f, err := os.Create(mainFile)
	if err!=nil{
		message.FormatError(func() {
			panic(fmt.Sprintf("Create main.go fail! error:%v",err.Error()))
		})
	}

	w := bufio.NewWriter(f) 
	_, err = w.WriteString(fmt.Sprintf(MAIN_FILE_TEMPLATE,params.ProjectName))
	if err!=nil{
		message.FormatError(func() {
			panic(fmt.Sprintf("Write main.go fail! error:%v",err.Error()))
		})
	}
	w.Flush()
	f.Close()
	
	message.Success(fmt.Sprintf("created %v",mainFile))
}

