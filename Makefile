PREFIX ?= ${GOPATH}

all: test

vendor:
	go mod vendor
	go mod tidy

test: vendor
	script/run_tests.sh

.PHONY: all test vendor
