BINARY = secrets
GOARCH = amd64

all: test clean linux darwin windows

clean:
	rm -rf bin/

linux:
	GOOS=linux GOARCH=${GOARCH} go build -o bin/${BINARY}-linux-${GOARCH}

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build -o bin/${BINARY}-darwin-${GOARCH}

windows:
	GOOS=windows GOARCH=${GOARCH} go build -o bin/${BINARY}-windows-${GOARCH}.exe

test:
	dep ensure
	go test ./...

install: test
	go install

.PHONY: all test clean linux darwin windows
