#!/bin/bash
set -e

export CGO_ENABLED=0

rm -rf goroot
mkdir goroot
ln -s $(go env GOROOT)/test goroot/test
ln -s $(go env GOROOT)/lib goroot/lib

cp -r $(go env GOROOT)/src goroot/src

mkdir goroot/pkg
ln -s $(go env GOROOT)/pkg/tool goroot/pkg/tool
ln -s $(go env GOROOT)/pkg/include goroot/pkg/include

PACKAGES=$(cd $(go env GOROOT)/src; go list ./... | egrep -v "runtime|builtin|cmd|sync")

go build rewrite_package.go
for pkg in $PACKAGES; do
	echo $pkg
	./rewrite_package $pkg
done

env GOROOT=$PWD/goroot $(which go) install -v $PACKAGES

env GOROOT=$PWD/goroot $(which go) test -short $PACKAGES
