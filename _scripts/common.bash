# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

set -u
set -o pipefail

# The following must be set _before_ the trap in order that the trap also applies within
# function bodies
#
# See https://www.gnu.org/software/bash/manual/html_node/Shell-Functions.html#Shell-Functions
set -o errtrace

shopt -s globstar
shopt -s extglob

error() {
  local lineno="$1"
  local file="$2"

  # intentional so we can test BASH_SOURCE
  if [[ -n "$file" ]] ; then
    echo "Error on line $file:$lineno"
  fi

  exit 1
}

trap 'set +u; error "${LINENO}" "${BASH_SOURCE}"' ERR
