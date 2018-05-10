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

{
	pushd examples/sites/helloworldbootstrap

	rm -f *.{go,html}
	reactGen -init bootstrap

	popd
}

gg ./...

z=$(goimports -l !(_vendor|_talks)/**/!(gen_*).go !(gen_*).go)
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

# with Go 1.10 we have to manually install deps of vetters below
# because package dependencies aren't automatically built into
# GOPATH/pkg. This will be fixed in later Go versions... by having
# the go command pass values to a vetter telling it (the vetter)
# where built packages exist

go list -f "{{ range .Deps}}{{.}}
{{end}}" ./... | xargs go install

go test ./...

gjbt myitcv.io/react

go vet ./...

reactVet ./...

immutableVet ./...

# we need to explicitly test the generated test files
go test myitcv.io/react/cmd/stateGen/_testFiles/

if [ "${CI:-}" == "true" ]
then
	# off the back of https://github.com/myitcv/react/issues/116#issuecomment-380280847
	# ensure that we can go get myitcv.io/react/... without _vendor
	echo "Verify that we can go get myitcv.io/react/..."
	export GOPATH=$HOME/gopath
	go get myitcv.io/react/...
fi
