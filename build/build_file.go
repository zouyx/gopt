package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
)


const (
	BUILD_FILE_TEMPLATE="#!/bin/bash\n" +
"\n" +
"fileName=\"package.properties\"\n" +
"\n" +
"build(){\n" +
"    read -p \"enter your GOPATH(no enter will use parent path of src): \" goPath\n" +
"    setEnv\n" +
"\n" +
"    echo \"installing package!\"\n" +
"\n" +
"    reimport\n" +
"\n" +
"    echo \"build successfully!\"\n" +
"}\n" +
"\n" +
"reimport() {\n" +
"    echo \"==============import==============\"\n" +
"    setEnv\n" +
"    echo \"Current GOPATH :\"\n" +
"    echo $GOPATH\n" +
"\n" +
"    for line in `cat $fileName`\n" +
"    do\n" +
"        echo \"installing $line!\"\n" +
"        go get \"$line\" -v -u\n" +
"    done\n" +
"}\n" +
"\n" +
"setEnv(){\n" +
"    echo \"==============Set path==============\"\n" +
"    if [ -n \"$goPath\" ]; then\n" +
"        export GOPATH=$goPath\n" +
"    else\n" +
"        currentPath=`pwd`\n" +
"        export GOPATH=${currentPath%src*}\n" +
"    fi\n" +
"}\n" +
"\n" +
"case \"$1\" in\n" +
"  build)\n" +
"\tbuild ;;\n" +
"  reimport)\n" +
"\treimport ;;\n" +
"  *)\n" +
"    echo $\"Usage: $0 {build|reimport}\"\n" +
"    exit 1\n" +
"esac\n"
	BUILD_FILE="/build.sh"
)

type BuildBuilder struct {

}

// build.sh build method
func (this *BuildBuilder) Build(params *input.Params) {
	src := getSrc(params)
	fileName := src + BUILD_FILE

	write(fileName,fmt.Sprintf(BUILD_FILE_TEMPLATE))
}
