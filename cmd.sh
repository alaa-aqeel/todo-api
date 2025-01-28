#!/bin/bash

case "$1" in
    "up")
        sudo docker compose -f docker-compose.yml up --build
        ;;
    "test")
        docker compose exec testing go test -v ./src/test/**/*
        ;;
    *)
        echo "Usage: $0 {up|test}"
        exit 1
        ;;
esac