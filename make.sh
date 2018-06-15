#!/bin/sh

bin_name=gopt

GOOS=linux   GOARCH=amd64 go build -o ./dist/linux_64_${bin_name}
GOOS=linux   GOARCH=386   go build -o ./dist/linux_386_${bin_name}
GOOS=windows GOARCH=386   go build -o ./dist/windows_386_${bin_name}.exe
GOOS=windows GOARCH=amd64 go build -o ./dist/windows_64_${bin_name}.exe
GOOS=darwin  GOARCH=386   go build -o ./dist/darwin_386_${bin_name}
GOOS=darwin  GOARCH=amd64 go build -o ./dist/darwin_64_${bin_name}

cd dist

mv  linux_386_${bin_name}       ${bin_name}
zip linux_386_${bin_name}       ${bin_name}     -qm
mv  linux_64_${bin_name}        ${bin_name}
zip linux_64_${bin_name}        ${bin_name}     -qm
mv  windows_386_${bin_name}.exe ${bin_name}.exe
zip windows_386_${bin_name}     ${bin_name}.exe -qm
mv  windows_64_${bin_name}.exe  ${bin_name}.exe
zip windows_64_${bin_name}      ${bin_name}.exe -qm
mv  darwin_386_${bin_name}      ${bin_name}
zip darwin_386_${bin_name}      ${bin_name}     -qm
mv  darwin_64_${bin_name}       ${bin_name}
zip darwin_64_${bin_name}       ${bin_name}     -qm

cd -