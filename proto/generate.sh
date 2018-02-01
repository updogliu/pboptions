#!/usr/bin/env bash

# Clean up the current generated protobuf go codes
find . -name \*.pb.go -delete

GOSRC="$GOPATH"/src

protoc -I="$GOSRC" \
 -I=. --go_out=plugins=grpc:"$GOSRC" \
 *.proto
