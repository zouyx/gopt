package input

import (
	"testing"
)

func TestWelcome(t *testing.T) {
	(&WelcomeHandler{}).Handle(nil)
}
