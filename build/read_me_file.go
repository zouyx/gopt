package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
)


const (
	READ_ME_FILE_TEMPLATE="# %v\n" +
		"Project template for new Golang project and integrate some CI Tools like Travis CI,coveralls etc.\n"
	READ_ME_FILE="/README.md"
)

type ReadMeBuilder struct {

}

// README.md build method
func (this *ReadMeBuilder) Build(params *input.Params) {
	src := getSrc(params)
	fileName := src + READ_ME_FILE

	write(fileName,fmt.Sprintf(READ_ME_FILE_TEMPLATE,params.ProjectName))
}
