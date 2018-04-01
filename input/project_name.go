package input

import (
	"fmt"
	"github.com/zouyx/gopt/message"
)

type ProjectNameInputHandler struct {
}

//init project name to Params with user input
func (this *ProjectNameInputHandler) Handle(params *Params) {
	fmt.Println("1.Choose your project name:")
	line, err := readLine()
	if len(line)==0||err!=nil{
		message.FormatError(func() {
			panic("cannot read your project name.")
		})
	}

	params.ProjectName=string(line)
}
