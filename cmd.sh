#!/bin/bash

case "$1" in
    "up")
        sudo docker compose -f docker-compose.yml up --build
        ;;
    "test")
        docker compose exec testing go test -v ./src/test/$2
        ;;
    *)
        echo "Usage: $0 {up|test [path]}"
        exit 1
        ;;
esac