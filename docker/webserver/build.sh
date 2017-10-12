#!/usr/bin/env bash
set -e
echo "buidling linux executable"
GOOS=linux go build
docker build -t huangjoyce3/testserver .
docker push huangjoyce3/testserver 
go clean
