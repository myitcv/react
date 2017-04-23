#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "${BASH_SOURCE%/*}/common.bash"

r=$(mktemp -d)
t=$(mktemp -d)

echo "Cloning git@github.com:myitcv/gopherjs_examples_sites.git into $r"

git clone -q git@github.com:myitcv/gopherjs_examples_sites.git $r/gopherjs_examples_sites
rm -rf $r/gopherjs_examples_sites/*

echo ""

echo "Copying..."

for i in $(command ls "${BASH_SOURCE%/*}/../examples/sites" | grep -v common)
do
	echo $i
	(
		cd $t
		wget --quiet --mirror http://localhost:8080/myitcv.io/react/examples/sites/$i/
	)
	cp -rp $t/localhost:8080/myitcv.io/react/examples/sites/$i/ $r/gopherjs_examples_sites/

	du -sh $r/gopherjs_examples_sites/$i
done

cp -rp examples/sites/common $r/gopherjs_examples_sites/

echo ""

cd $r/gopherjs_examples_sites
git config hooks.stopbinaries false

if [ -z "$(git status --porcelain)" ]
then
	echo "No changes to commit"
	exit 0
fi

git add -A
git commit -am "Examples update at $(date)"
git push
