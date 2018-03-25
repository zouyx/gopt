package main

import (
	"fmt"
	"github.com/zouyx/gopt/input"
	"encoding/json"
)

func main() {
	param := input.GetAllParams()
	bytes, _ := json.Marshal(param)
	fmt.Println("hello world!",string(bytes))
}
