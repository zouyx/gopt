package main

import (
	"fmt"
	"github.com/zouyx/gopt/input"
	"encoding/json"
)

func main() {
	param := input.GetParam()
	bytes, _ := json.Marshal(param)
	fmt.Println("hello world!",string(bytes))
}
