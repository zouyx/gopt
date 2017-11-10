read -p "enter your GOPATH: " goPath
if [ -n "$goPath" ]; then
    export GOPATH=$goPath
    echo "GOPATH has change to $goPath"
fi

echo "installing seelog!"

gopm get github.com/cihub/seelog -v -g

echo "build successfully!"