PREFIX ?= ${GOPATH}

all: test

vendor:
	go mod vendor

test: vendor
	script/run_tests.sh

.PHONY: all test vendor
