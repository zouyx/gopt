package message

import "fmt"

func Success(message string)  {
	fmt.Println("gopt : ",message)
}

// format error for print
func FormatError(printError func())  {
	fmt.Printf("gopt : there is something wrong ! ")
	printError()
	fmt.Println()
}