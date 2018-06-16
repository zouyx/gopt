package message

import (
	"testing"
	"fmt"
)

func TestFormatError(t *testing.T) {
	FormatError(func() {
		fmt.Sprintf("Create %v fail! error:%v","file","error!")
	})
}

func TestSuccess(t *testing.T) {
	Success("hello!")
}
