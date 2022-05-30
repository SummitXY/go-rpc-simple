#!/bin/bash

RUN_NAME="go-rpc-simple"

mkdir -p output
go build -a -o output/bin/${RUN_NAME}
chmod +x output/bin/${RUN_NAME}

