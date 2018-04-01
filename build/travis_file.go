package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
)


const (
	TRAVIS_FILE_TEMPLATE="language: go\n" +
		"sudo: required\n" +
		"\n" +
		"go:\n" +
		"  - 1.7\n" +
		"  - 1.8\n" +
		"  - 1.9\n" +
		"\n" +
		"env:\n" +
		"    - GO15VENDOREXPERIMENT=\"1\"\n" +
		"\n" +
		"before_install:\n" +
		"    - go get github.com/mattn/goveralls\n" +
		"\n" +
		"script:\n" +
		"    - chmod u+x coverage.sh\n" +
		"    - ./coverage.sh --coveralls\n" +
		"    - go build -o \"gopt\" main/main.go"
	TRAVIS_FILE="/.travis.yml"
)

type TravisBuilder struct {

}

// .travis.yml build method
func (this *TravisBuilder) Build(params *input.Params) {
	src := getSrc(params)
	fileName := src + TRAVIS_FILE

	write(fileName,fmt.Sprintf(TRAVIS_FILE_TEMPLATE))
}
