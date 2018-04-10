#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "${BASH_SOURCE%/*}/common.bash"

# ensure we are in the right directory
cd "${BASH_SOURCE%/*}/.."

export PATH=$PWD/_vendor/bin:$GOPATH/bin:$PATH
export GOPATH=$PWD/_vendor:$GOPATH

go install myitcv.io/react/cmd/reactGen

go generate
