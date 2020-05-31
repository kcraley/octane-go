
# Variables used for compiling the octane binary
binary=octane
module=github.com/kcraley/octane-go
buildDate=$(shell date "+%s")
gitVersion=$(shell git describe --tags --always)

LDFLAGS=\
  -X $(module)/version.buildDate=$(buildDate) \
  -X $(module)/version.gitVersion=$(gitVersion)


deps:
	go get

build: deps
	go build -ldflags="$(LDFLAGS)" -o $(binary) .
