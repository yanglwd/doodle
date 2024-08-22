#!/bin/bash
set -e

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
if [ $? -ne 0 ]; then
    echo "\033[0;33m protoc-gen-go install failed. Aborting.\033[m"
    exit 1
else
    echo "protoc-gen-go install success. Continue."
fi

protoc --proto_path=. --go_out=../message/ *.proto