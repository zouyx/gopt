package main

import (
	"fmt"
	"encoding/json"
	"github.com/zouyx/gopt/param"
)

func main() {
	params := param.GetAllParams()
	bytes, _ := json.Marshal(params)
	fmt.Println("hello world!",string(bytes))
}
