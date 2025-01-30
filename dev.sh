#!/bin/bash

# might need host_mnt (got issue using docker desktop on linux only weird??)
# export CODE_ROOT="/host_mnt${PWD}"
export CODE_ROOT="${PWD}"


echo "CODE_ROOT: $CODE_ROOT"

docker build -t mesh-demo/frontend:dev -f frontend/deployment/Dockerfile.dev frontend/
docker build -t mesh-demo/backend:dev -f backend/deployment/Dockerfile.dev backend/
docker build -t mesh-demo/llm:dev -f llm/deployment/Dockerfile.dev llm/

envsubst < frontend/deployment/deploy.dev.yaml | kubectl apply -f -
envsubst < backend/deployment/deploy.dev.yaml | kubectl apply -f -
envsubst < llm/deployment/deploy.dev.yaml | kubectl apply -f -
