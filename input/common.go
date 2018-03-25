package input

import (
	"bufio"
	"os"
)

// read line method
func readLine() (string,error) {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	return string(line),err
}


type Params struct {
	/**
	project type
	0:custom
	1:normal
	2:github
	 */
	ProjectType int8 `json:"project_type"`
}