
MACHINE_NAME=docker

DOCKER=eval $$(docker-machine env $(MACHINE_NAME)) && docker
COMPOSE=eval $$(docker-machine env $(MACHINE_NAME)) && docker-compose --x-networking -p test

HOST=$(shell docker-machine ip $(MACHINE_NAME))

all: deps build up

deps:
	go get ./...

build:
	$(COMPOSE) build

up:
	$(COMPOSE) up -d

status:
	$(COMPOSE) ps

logs:
	$(COMPOSE) logs httpd

browser:
	xdg-open http://$(HOST):80

