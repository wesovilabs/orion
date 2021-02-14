#!/bin/bash
echo "publishing docker image..."
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push wesovilabs/orion
