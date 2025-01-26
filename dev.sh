#!/bin/bash

# For Linux
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    export CODE_ROOT="/host_mnt${PWD}"
# For macOS
elif [[ "$OSTYPE" == "darwin"* ]]; then
    export CODE_ROOT="${PWD}"
fi
# might need to add for windows

echo "CODE_ROOT: $CODE_ROOT"

envsubst < frontend/deployment/deploy.dev.yaml | kubectl apply -f -
envsubst < backend/deployment/deploy.dev.yaml | kubectl apply -f -
