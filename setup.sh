#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

RETVAL=0

go get github.com/appscode/api/...
pushd ${GOPATH}/src/github.com/appscode/api
./gen
popd
go install
exit ${RETVAL}