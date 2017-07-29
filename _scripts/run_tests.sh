#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "${BASH_SOURCE%/*}/common.bash"

export PATH=$PWD/_vendor/bin:$GOPATH/bin:$PATH
export GOPATH=$PWD/_vendor:$GOPATH

# ensure we are in the right directory
cd "${BASH_SOURCE%/*}/.."

for i in $(cat .vendored_bin_deps .bin_deps)
do
	go install $i
done

find -path ./_vendor -prune -o -name "gen_*.go" -exec rm '{}' \;

{
	pushd examples/sites/helloworld

	rm -f *.{go,html}
	reactGen -init minimal

	popd
}

go generate ./...

z=$(goimports -l !(_vendor)/**/!(gen_*).go !(gen_*).go)
if [ ! -z "$z" ]
then
	echo "The following files are not formatted:"
	echo ""
	echo "$z"
	exit 1
fi

z=$(gofmt -l !(_vendor)/**/gen_*.go gen_*.go)

if [ ! -z "$z" ]
then
	echo "The following generated files are not formatted:"
	echo ""
	echo "$z"
	exit 1
fi

# we need to install first so the go/types-based reactVet tests
# can import the myitcv.io/react/jsx package
go install ./...

go test ./...

go vet ./...

reactVet ./...

immutableVet ./...

# we need to explicitly test the generated test files
go test myitcv.io/react/cmd/stateGen/_testFiles/

