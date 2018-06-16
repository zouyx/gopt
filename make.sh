#!/bin/sh

bin_name=gopt

GOOS=linux   GOARCH=amd64 go build -o ./dist/linux_64_${bin_name}
GOOS=linux   GOARCH=386   go build -o ./dist/linux_386_${bin_name}
GOOS=windows GOARCH=386   go build -o ./dist/windows_386_${bin_name}.exe
GOOS=windows GOARCH=amd64 go build -o ./dist/windows_64_${bin_name}.exe
GOOS=darwin  GOARCH=386   go build -o ./dist/darwin_386_${bin_name}
GOOS=darwin  GOARCH=amd64 go build -o ./dist/darwin_64_${bin_name}

cd dist

pwd

cd -