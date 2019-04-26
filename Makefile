APPNAME = kpaas
BIN = $(GOPATH)/bin
GOCMD = /usr/local/go/bin/go
GOBUILD = $(GOCMD) build
GOINSTALL = $(GOCMD) install
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GORUN = $(GOCMD) run
BINARY_UNIX = $(BIN)_unix

test:
	$(GOTEST) -v ./...

install:
	curl https://glide.sh/get | sh
	glide install

run:
	$(GORUN) ./cmd/server/server.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v