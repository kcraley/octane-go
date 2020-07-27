
# Variables used for compiling the octane binary
binary=octane
module=github.com/kcraley/octane-go
buildDate=$(shell date "+%s")
gitVersion=$(shell git describe --tags --always)

LDFLAGS=\
  -X $(module)/version.buildDate=$(buildDate) \
  -X $(module)/version.gitVersion=$(gitVersion)


.PHONY: container
container: Dockerfile
	docker build -t $(binary):$(gitVersion) -f $< .

.PHONY: serve
serve: container
	docker run -it --rm --name $(binary) \
	-e COD_USERNAME=$(cod-username) \
	-e COD_PASSWORD=$(cod-password) \
	$(binary):$(gitVersion) \
	serve -t $(token)

.PHONY: deps
deps:
	go get

.PHONY: build
build: deps
	go build -ldflags="$(LDFLAGS)" -o $(binary) .

.PHONY: test
test: deps
	go test -v ./...

.PHONY: vet
vet: deps
	go vet
