package build

import (
	"github.com/zouyx/gopt/input"
	"github.com/zouyx/gopt/message"
	"fmt"
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

	write(mainFile,fmt.Sprintf(MAIN_FILE_TEMPLATE,params.ProjectName))

	message.Success(fmt.Sprintf("Created %v",mainFile))
}

