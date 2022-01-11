#! /bin/bash
GOOS=linux GOARCH=amd64 go build -o api-linux main.go
# docker image build -t grpc-api .
# docker container run -p 50051:50051 grpc-api
