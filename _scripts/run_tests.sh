#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "${BASH_SOURCE%/*}/common.bash"

export PATH=$PWD/_vendor/bin:$GOPATH/bin:$PATH
export GOPATH=$PWD/_vendor:$GOPATH

for i in $(cat .vendored_bin_deps .bin_deps)
do
	go install $i
done

go generate ./...

go test ./...

go install ./...

go vet ./...

immutableVet ./...
