#!/bin/bash

unset HTTP_PROXY
unset HTTPS_PROXY

DOCKER_COMPOSE_FILE="docker-compose.yml"
DOCKER_COMPOSE_CONTAINER_NAME="local_rgo001"
DATABASE_NAME="rgo001_users_active_delivery"
USER_NAME="rgo001_delivery_orchestror"
PORT="3035"

# Função para validar a execução do Docker
if ! docker info > /dev/null; then
  echo "Docker is not running. Please start Docker and try again."
  exit 1
else

  # Criação de Container
  if docker container inspect $DOCKER_COMPOSE_CONTAINER_NAME > /dev/null; then
    echo "Container $DOCKER_COMPOSE_CONTAINER_NAME already exists. Skipping..."
  else
    echo "Creating container..."
    docker-compose -f $DOCKER_COMPOSE_FILE --project-name $DOCKER_COMPOSE_CONTAINER_NAME up -d
    sleep 5

    if [ $? -eq 0 ]; then
      echo "Container $DOCKER_COMPOSE_CONTAINER_NAME done!"
    fi
  fi
fi

#Inicializar bando de dados via node server
cd db-init-scripts/node
npm install
node server.js &

echo 'Initializing database with mongoDBexpress & node'

curl http://localhost:$PORT/create-database
sleep 5

echo 'Database done. All Steps Done. Ty for waiting!'