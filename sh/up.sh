#!/bin/bash

echo 'Starting docker compose'
docker compose up -d

echo 'Connecting to container'
docker exec -it uploader bash

