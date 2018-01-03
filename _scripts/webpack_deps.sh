#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "${BASH_SOURCE%/*}/common.bash"

pushd internal/dev > /dev/null
{
	npm install
}
popd > /dev/null

for i in dev prod testutils
do
	pushd internal/$i > /dev/null
	{
		./node_modules/.bin/webpack
	}
	popd > /dev/null
done
