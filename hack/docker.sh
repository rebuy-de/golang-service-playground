#/bin/bash

source $( dirname $0)/env.sh

hack/deps.sh

docker-compose --x-networking build
docker-compose --x-networking up --rm "$@"
