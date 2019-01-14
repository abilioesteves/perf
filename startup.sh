#!/bin/bash
set -e
set -x

echo "Starting the container performance checker..."
./perf \
    --wp=${WRITE_PATH}