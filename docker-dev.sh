#!/bin/sh -ex

docker network create -d bridge damnswank || true

docker kill damnswank-dev || true
sleep 2

docker build -t damnswank-dev \
	     -f Dockerfile \
	     --build-arg "DATE=$(date)" \
	     --build-arg "REVISION=$(git rev-parse HEAD)" \
	     .
sleep 2
docker run --rm --name damnswank-dev -d --network=damnswank -p 8002:8002 damnswank-dev
sleep 5
docker ps
echo "docker run --rm --name damnswank-dev -ti -p 8002:8002 damnswank-dev bash"
docker logs damnswank-dev
echo "docker exec -ti damnswank-dev bash"
