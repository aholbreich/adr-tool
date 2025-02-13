BINARY_NAME=adr
INSTALL_DIR=$(HOME)/bin

build:
	go fmt
	go mod tidy
	go build -o $(BINARY_NAME)

test: build
	go test -v ./...

install: build test
	mkdir -p $(INSTALL_DIR)
	mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

cleancache:
	go clean -cache -testcache -modcache
