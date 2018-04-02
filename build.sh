#!/bin/bash

fileName="package.properties"

build(){
    read -p "enter your GOPATH(no enter will use parent path of src): " goPath
    setEnv

    echo "installing package!"

    reimport

    echo "build successfully!"
}

run(){
    read -p "enter your GOPATH(no enter will use parent path of src): " goPath
    setEnv

    echo "running...."
    go run main/main.go

}

reimport() {
    echo "==============import=============="
    setEnv
    echo "Current GOPATH :"
    echo $GOPATH

    for line in `cat $fileName`
    do
        echo "installing $line!"
        go get "$line" -v -u
    done
}

setEnv(){
    echo "==============Set path=============="
    if [ -n "$goPath" ]; then
        export GOPATH=$goPath
    else
        currentPath=`pwd`
        export GOPATH=${currentPath%src*}
    fi
}

case "$1" in
  build)
	build ;;
  reimport)
	reimport ;;
  run)
	run ;;
  *)
    echo $"Usage: $0 {build|reimport|run}"
    exit 1
esac
