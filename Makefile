BINARY_NAME=cpu_monitor
VERSION=1.0.0

build:
	go build -o $(BINARY_NAME) ./src

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux ./src

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe ./src

build-all: build build-linux build-windows

install:
	go install ./src

clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME)-linux $(BINARY_NAME).exe

test:
	go test ./src

.PHONY: build install clean test
