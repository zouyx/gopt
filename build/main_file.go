package build

import (
	"github.com/zouyx/gopt/input"
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
		MAIN_FILE="/main.go"
)


type MainFileBuilder struct {

}

// main.go build method
func (this *MainFileBuilder) Build(params *input.Params) {
	fullPath := getFullPath(params)
	fileName := fullPath + MAIN_FILE

	write(fileName,fmt.Sprintf(MAIN_FILE_TEMPLATE,params.ProjectName))
}

