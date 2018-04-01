package input

import "fmt"


type WelcomeHandler struct {
}

//print welcome message
func (this *WelcomeHandler) Handle(params *Params) {
	fmt.Println("===============gopt=============")
	fmt.Println("Thanks you to use gopt!")
	fmt.Println("===============gopt=============")
}
