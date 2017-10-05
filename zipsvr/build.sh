#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t huangjoyce3/zipsvr .
docker push huangjoyce3/zipsvr
go clean 
