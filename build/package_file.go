package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
)


const (
	PACKAGE_FILE_TEMPLATE="github.com/zouyx/gomatcher"
	PACKAGE_FILE="/package.properties"
)

type PackageBuilder struct {

}

// package.properties build method
func (this *PackageBuilder) Build(params *input.Params) {
	src := getSrc(params)
	fileName := src + PACKAGE_FILE

	write(fileName,fmt.Sprintf(PACKAGE_FILE_TEMPLATE))
}
