#!/bin/bash
if [ "x$(uname)" == "xDarwin" ]; then
    BIN_DIR=$(greadlink -f $(dirname $0))
else
    BIN_DIR=$(readlink -f $(dirname $0))
fi
ROOT_DIR=`dirname $BIN_DIR`

#$BIN_DIR/project-layout-go -p $ROOT_DIR -c conf/app.yaml
$BIN_DIR/project-layout-go
