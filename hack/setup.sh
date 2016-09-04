#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

RETVAL=0

deps() {
    go get github.com/appscode/api/...
    local -r owd=$PWD
    cd $GOPATH/src/github.com/appscode/api
    ./hack/gen.sh
    cd $owd
}

godep() {
    local -r owd=$PWD
    cd $GOPATH/src/github.com/appscode/client
    rm -rf Godeps vendor
    GO15VENDOREXPERIMENT=1 $GOPATH/bin/godep save ./...
    cd $owd
}

build() {
    local -r owd=$PWD
    cd $GOPATH/src/github.com/appscode/client
    go get github.com/jteeuwen/go-bindata/...
    go-bindata -ignore=\\.go -o files/data.go -pkg files files/...
    go build ./...
    cd $owd
}

if [ $# -eq 0 ]; then
	build
	exit $RETVAL
fi

case "$1" in
	deps)
		deps
		;;
	godep)
		godep
		;;
	build)
		build
		;;
	*)	(10)
		echo $"Usage: $0 {deps|godep|build}"
		RETVAL=1
esac
exit $RETVAL
