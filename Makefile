# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate

# App parameters
GOPI=github.com/djthorpe/gopi
GOLDFLAGS += -X $(GOPI).GitTag=$(shell git describe --tags)
GOLDFLAGS += -X $(GOPI).GitBranch=$(shell git name-rev HEAD --name-only --always)
GOLDFLAGS += -X $(GOPI).GitHash=$(shell git rev-parse HEAD)
GOLDFLAGS += -X $(GOPI).GoBuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GOFLAGS = -ldflags "-s -w $(GOLDFLAGS)" 

all: test install

install: helloworld-client helloworld-service discovery-service discovery-client dns-discovery

protobuf:
	$(GOGEN) -x ./rpc/...

helloworld-client: protobuf
	$(GOINSTALL) $(GOFLAGS) ./cmd/helloworld-client/...

helloworld-service: protobuf
	$(GOINSTALL) $(GOFLAGS) ./cmd/helloworld-service/...

discovery-service: protobuf
	$(GOINSTALL) $(GOFLAGS) ./cmd/discovery-service/...

discovery-client: protobuf
	$(GOINSTALL) $(GOFLAGS) ./cmd/discovery-client/...

dns-discovery:
	$(GOINSTALL) $(GOFLAGS) ./cmd/dns-discovery/...

test:  protobuf
	$(GOTEST) ./...

clean: 
	$(GOCLEAN)

