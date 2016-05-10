#/bin/bash

source $( dirname $0)/env.sh

hack/build.sh

cd target

rm -rf ./usr/
mkdir -p ./usr/bin/

cp ${project_name} ./usr/bin/

fpm \
	-s dir \
	-t rpm \
	-n "rebuy-${project_name}" \
	-v ${project_version} \
	./usr/bin/${project_name}
