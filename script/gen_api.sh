#!/bin/bash

# Different working directory generates different results.
# Not sure if it's required to do this:

# Got these lines from: https://grpc.io/docs/languages/go/quickstart/
cd api

protoc \
--go_out=./contactpb \
--go_opt=paths=source_relative \
--go-grpc_out=./contactpb \
--go-grpc_opt=paths=source_relative \
-I. \
-I$(go env GOPATH)/src/github.com/googleapis/googleapis \
./contact_message.proto

protoc \
--go_out=./contactpb \
--go_opt=paths=source_relative \
--go-grpc_out=./contactpb \
--go-grpc_opt=paths=source_relative \
-I. \
-I$(go env GOPATH)/src/github.com/googleapis/googleapis \
./contact.proto

protoc \
--go_out=./utilpb \
--go_opt=paths=source_relative \
--go-grpc_out=./utilpb \
--go-grpc_opt=paths=source_relative \
-I. \
-I$(go env GOPATH)/src/github.com/googleapis/googleapis \
./util.proto
