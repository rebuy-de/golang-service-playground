#!/bin/bash

source $( dirname $0)/env.sh

hack/deps.sh

go test -p 1 $(hack/glidew.sh novendor)
