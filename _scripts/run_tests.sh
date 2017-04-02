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

rm -f !(_vendor)/**/gen_*.go

{
	pushd sites/helloworld

	rm -f *.{go,html}
	reactGen -init minimal

	popd
}

go generate ./...

go test ./...

go install ./...

go vet ./...

immutableVet ./...

# we need to explicitly test the generated test files
go test github.com/myitcv/gopherjs/cmd/stateGen/_testFiles/

