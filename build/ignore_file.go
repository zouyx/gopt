package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
	"github.com/zouyx/gopt/message"
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
)

type IgnoreBuilder struct {

}

// dir build method
func (this *IgnoreBuilder) Build(params *input.Params) {
	src := getSrc(params)
	gitIgnore := src + "/.gitignore"

	write(gitIgnore,fmt.Sprintf(IGNORE_FILE_TEMPLATE))

	message.Success(fmt.Sprintf("Created %v",gitIgnore))
}
