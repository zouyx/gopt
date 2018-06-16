#!/bin/sh

bin_name=gopt

echo "making linux_64..."
GOOS=linux   GOARCH=amd64 go build -o ./dist/${bin_name}_linux_64 main/main.go

echo "making linux_386..."
GOOS=linux   GOARCH=386   go build -o ./dist/${bin_name}_linux_386 main/main.go

echo "making windows_386..."
GOOS=windows GOARCH=386   go build -o ./dist/${bin_name}_windows_386.exe main/main.go

echo "making windows_64..."
GOOS=windows GOARCH=amd64 go build -o ./dist/${bin_name}_windows_64.exe main/main.go

echo "making darwin_386..."
GOOS=darwin  GOARCH=386   go build -o ./dist/${bin_name}_darwin_386 main/main.go

echo "making darwin_64..."
GOOS=darwin  GOARCH=amd64 go build -o ./dist/${bin_name}_darwin_64 main/main.go

cp dist/* .