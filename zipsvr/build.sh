#!/usr/bin/env bash
set -e
echo "building go server for Linux..."
GOOS=linux go build
docker build -t huangjoyce3/zipsvr .
docker push huangjoyce3/zipsvr
go clean 

