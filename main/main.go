package main

import (
	"fmt"
	"encoding/json"
	"github.com/zouyx/gopt/component"
)

func main() {
	//get params
	params := component.GetAllParams()
	bytes, _ := json.Marshal(params)

	//build
	component.BuildProject(params)
	fmt.Println("hello world!",string(bytes))
}
