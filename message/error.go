package message

import "fmt"

func FormatError(printError func())  {
	fmt.Printf("gopt : there is something wrong!")
	printError()
	fmt.Println()
}