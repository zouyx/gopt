package main

import (
	"fmt"
	"encoding/json"
	"github.com/zouyx/gopt/component"
	"github.com/zouyx/gopt/message"
)

func main() {
	//get params
	params := component.GetAllParams()

	//build
	component.BuildProject(params)

	bytes, _ := json.Marshal(params)
	message.Success(fmt.Sprintf("Finished make new project!your params are:%v",string(bytes)))
}
