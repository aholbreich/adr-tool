# Makefile
BINARY_NAME=adr
INSTALL_DIR=$(HOME)/bin
VERSION=$(shell git describe --tags --abbrev=0)
COMMIT_HASH=$(shell git rev-parse --short HEAD)
COUNT=$(shell git rev-list $(VERSION)..HEAD --count)

# Build the binary
build:
	go fmt
	go mod tidy
	go build -o $(BINARY_NAME) -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"

# Run tests
test: build
	go test -v ./...

# Install the binary
install: build test
	mkdir -p $(INSTALL_DIR)
	mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Clean Go build caches
cleancache:
	go clean -cache -testcache -modcache
