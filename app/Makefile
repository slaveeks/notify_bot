BINARY_NAME=notify

export GO111MODULE=on

all: check lint build

build:
	go build -o $(BINARY_NAME) -v ./
	chmod +x $(BINARY_NAME)
test:
	go test ./...
lint:
	golint ./
clean:
	go clean
	rm -rf $(BINARY_NAME)
	

build-all: build-linux

build-linux:
	GOOS=linux GOARCH=amd64 go build -o .$(BINARY_NAME_LINUX) -v ./
