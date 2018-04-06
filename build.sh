#!/bin/bash

docker build -t abc-server -f server.Dockerfile .
docker build -t abc-client -f client.Dockerfile .
