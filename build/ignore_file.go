package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
)


const (
	IGNORE_FILE_TEMPLATE="# Binaries for programs and plugins\n" +
		"*.exe\n" +
		"*.dll\n" +
		"*.so\n" +
		"*.dylib\n" +
		"\n" +
		"# Test binary, build with `go test -c`\n" +
		"*.test\n" +
		"\n" +
		"# Output of the go coverage tool, specifically when used with LiteIDE\n" +
		"*.out"

	IGNORE_FILE="/.gitignore"
)

type IgnoreBuilder struct {

}

// .gitignore build method
func (this *IgnoreBuilder) Build(params *input.Params) {
	src := getSrc(params)

	fileName := src + IGNORE_FILE

	write(fileName,fmt.Sprintf(IGNORE_FILE_TEMPLATE))
}
