package message

import "fmt"

// format error for print
func FormatError(printError func())  {
	fmt.Printf("gopt : there is something wrong ! ")
	printError()
	fmt.Println()
}