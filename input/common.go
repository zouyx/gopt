package input

import (
	"bufio"
	"os"
)

/**
read line method
 */
func readLine() (string,error) {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	return string(line),err
}