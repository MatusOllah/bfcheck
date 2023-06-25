GO = go

BINARY = ./bin

FLAGS = -gcflags="-dwarf=false" -v

GOOS = $(shell $(GO) env GOOS)
GOARCH = $(shell $(GO) env GOARCH)

.PHONY: all
all: test build

.PHONY: build
build: clean
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)
	$(GO) get
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

.PHONY: test
test:
	GO111MODULE=on $(GO) test -v

.PHONY: clean
clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
