# Command Definitions
GO          = go
GO_BUILD    = $(GO) build
GO_TEST     = $(GO) test -v

# Target Parameters
GO_PKGROOT  = ./

### API server build
.PHONY: builds
builds:
	$(GO_BUILD) $(GO_PKGROOT)...

### API server all test
.PHONY: tests
tests:
	$(GO_TEST) $(GO_PKGROOT)...

### API server package test
.PHONY: test
ifeq (test,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif
test:
	$(GO_TEST) $(GO_PKGROOT)/$(RUN_ARGS)

### GraphQL server build
.PHONY: generate
generate:
	gqlgen generate

### DI Generate
.PHONY: di
di:
	wire ./di/wire.go