#/bin/bash

source $( dirname $0)/env.sh

hack/test.sh

mkdir -p target
go build \
	-ldflags "-X main.version=${project_version}" \
	-o target/${project_name}
