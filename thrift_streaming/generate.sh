#!/bin/bash

cd `dirname $0`
DIR=`pwd`

set -e

module='-module github.com/cloudwego/kitex-tests'
idl=idl/api.thrift

SAVE_PATH=$PATH

# generate with old kitex and thriftgo WITHOUT thrift streaming support
function generate_old() {
    OLD=~/thrift-streaming/bin/github-old
    if [ ! -d "$OLD" ]; then
        mkdir -p $OLD
        # Put old kitex && thriftgo in ~/thrift-streaming/bin/github-old
        GOBIN=$OLD go install github.com/cloudwego/kitex/tool/cmd/kitex@v0.8.0
        GOBIN=$OLD go install github.com/cloudwego/thriftgo@v0.3.4
    fi
    export PATH=$OLD:$SAVE_PATH
    set -x
    kitex -gen-path kitex_gen_old $module $idl
}

function generate_new() {
    NEW=~/thrift-streaming/bin/github-new
    if [ ! -d "$NEW" ]; then
        echo "Please put new kitex/thriftgo supporting thrift streaming in ~/thrift-streaming/bin/github-new"
        exit 1
    fi
    export PATH=$NEW:$SAVE_PATH
    set -x

    # Thrift
    kitex $module $idl
    kitex -thrift template=slim -gen-path kitex_gen_slim $module $idl

    # KitexPB
    kitex $module idl/api.proto

    # GRPC
    kitex $module idl/api_no_stream.proto
}

# generate with old kitex WITHOUT thrift streaming support; no need to regenerate if no update in api.thrift
generate_old
generate_new