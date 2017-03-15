#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "${BASH_SOURCE%/*}/common.bash"

t=$(mktemp -d)

for i in $(command ls "${BASH_SOURCE%/*}/../sites")
do
	echo $i
	aws s3 rm --quiet s3://github.com.myitcv.gopherjs.react.examples/$i --recursive
	(
		cd $t
		wget --quiet -p -k http://localhost:8080/github.com/myitcv/gopherjs/sites/$i/
	)
	aws s3 cp --quiet $t/localhost:8080/github.com/myitcv/gopherjs/sites/$i/ s3://github.com.myitcv.gopherjs.react.examples/$i --recursive
done

