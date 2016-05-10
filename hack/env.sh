#!/bin/bash

project_version=$( 
	git describe --always --dirty \
		| tr '-' '.'
)
project_name=golang-service-playground

cd $( dirname $0 )/..
set -ex
