#!/bin/sh -ex

docker network create -d bridge damnswank || true

docker build -t damnswank-dev \
	     -f Dockerfile \
	     --build-arg "DATE=$(date)" \
	     --build-arg "REVISION=$(git rev-parse HEAD)" \
	     .

docker kill damnswank-dev || true
sleep 2

docker run --rm --name damnswank-dev -d --network=damnswank -p 8002:8002 damnswank-dev
sleep 5
docker ps

docker logs damnswank-dev
echo "docker run --rm --name damnswank-dev -ti -p 8002:8002 damnswank-dev bash"
