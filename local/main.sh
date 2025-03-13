#!/bin/bash

unset HTTP_PROXY
unset HTTPS_PROXY

DOCKER_COMPOSE_FILE="docker-compose.yml"
DOCKER_COMPOSE_CONTAINER_NAME="local_rgo001"

# Function to validate Docker Execution
if ! docker info > /dev/null; then
  echo "Docker is not running. Please start Docker and try again."
  exit 1
else
  # Docker Execution
  if docker container inspect $DOCKER_COMPOSE_CONTAINER_NAME > /dev/null; then
    echo "Container $DOCKER_COMPOSE_CONTAINER_NAME already exists. Skipping..."
  else
    docker-compose -f $DOCKER_COMPOSE_FILE --project-name $DOCKER_COMPOSE_CONTAINER_NAME up -d
    echo "Creating container..."
    sleep 10

    if [ $? -eq 0 ]; then
      echo "Container $DOCKER_COMPOSE_CONTAINER_NAME started successfully."
    fi
  fi
fi

echo "Ty for waiting!"