# Variables
PROJECT = gqlgen-todos
COMPOSE_PATH := ./docker-compose.yml
COMPOSE_COMMAND := docker-compose -f $(COMPOSE_PATH) -p $(PROJECT)
SERVICE_API := api

### Start selected environment
.PHONY: start
start:
	${COMPOSE_COMMAND} up -d --build

### Start selected environment
.PHONY: build
build:
	${COMPOSE_COMMAND} up -d --build ${SERVICE_API}

### Build and Start selected environment
.PHONY: restart
ifeq (logs,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif
restart:
	${COMPOSE_COMMAND} kill $(RUN_ARGS) \
 && ${COMPOSE_COMMAND} rm -f $(RUN_ARGS) \
 && ${COMPOSE_COMMAND} up -d --build $(RUN_ARGS)

### Shut down selected environment
.PHONY: stop
stop:
	${COMPOSE_COMMAND} kill \
 && ${COMPOSE_COMMAND} rm -f \

### Dell selected environment
.PHONY: down
ifeq (exec,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif
down:
	${COMPOSE_COMMAND} down $(RUN_ARGS)

### Up selected environment
.PHONY: up
up:
	${COMPOSE_COMMAND} up

 ## List containers
.PHONY: ps
ps:
	${COMPOSE_COMMAND} ps

### Show logs from selected environment
.PHONY: logs
ifeq (logs,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif
logs:
	${COMPOSE_COMMAND} logs -f $(RUN_ARGS)

### Enter Container
.PHONY: exec
ifeq (exec,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif
exec:
	${COMPOSE_COMMAND} exec $(RUN_ARGS) bash

### test
.PHONY: test-unit
test-unit:
	${COMPOSE_COMMAND} exec \
	-e GO_ENV=test \
	api richgo test -v ./src/.../.