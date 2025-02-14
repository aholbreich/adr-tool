# Makefile
BINARY_NAME=adr
INSTALL_DIR=$(HOME)/bin
VERSION=$(shell git describe --tags --abbrev=0)
COMMIT_HASH=$(shell git rev-parse --short HEAD)
COUNT=$(shell git rev-list $(VERSION)..HEAD --count)
# Define the target platforms
PLATFORMS=linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

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


get-version:
	@echo $(VERSION)-$(COUNT)-$(COMMIT_HASH)

bump:
	@echo "Current version: $(VERSION)"
	@echo "Current commit hash: $(COMMIT_HASH)"
	@echo "Current count: $(COUNT)"
	@echo "Enter new version: "
	@read new_version; \
	git tag $$new_version; \
	git push origin $$new_version

amend:
	git add .
	git commit --amend --no-edit
	git push --force


# Binaries distribution
binary_linux_amd64: 
	mkdir -p build/linux-amd64
	GOOS=linux GOARCH=amd64 go build -o build/linux-amd64/$(BINARY_NAME) -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"
	tar -C build/linux-amd64 -czvf adr-linux-amd64.tar.gz $(BINARY_NAME)
	rm -rf build/linux-amd64

binary_linux_arm64:
	mkdir -p build/linux-arm64
	GOOS=linux GOARCH=arm64 go build -o build/linux-arm64/$(BINARY_NAME) -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"
	tar -C build/linux-arm64 -czvf adr-linux-arm64.tar.gz $(BINARY_NAME)
	rm -rf build/linux-arm64

binary_darwin_amd64: 
	mkdir -p build/darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o build/darwin-amd64/$(BINARY_NAME) -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"
	tar -C build/darwin-amd64 -czvf adr-darwin-amd64.tar.gz $(BINARY_NAME)
	rm -rf build/darwin-amd64

binary_darwin_arm64: 
	mkdir -p build/darwin-arm64
	GOOS=darwin GOARCH=arm64 go build -o build/darwin-arm64/$(BINARY_NAME) -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"
	tar -C build/darwin-arm64 -czvf adr-darwin-arm64.tar.gz $(BINARY_NAME)
	rm -rf build/darwin-arm64

binary_windows_amd64: 
	mkdir -p build/windows-amd64
	GOOS=windows GOARCH=amd64 go build -o build/windows-amd64/$(BINARY_NAME).exe -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"
	zip -j adr-windows-amd64.zip build/windows-amd64/$(BINARY_NAME).exe
	rm -rf build/windows-amd64

binary_windows_arm64:
	mkdir -p build/windows-arm64
	GOOS=windows GOARCH=arm64 go build -o build/windows-arm64/$(BINARY_NAME).exe -ldflags "-X main.version=$(VERSION)-$(COUNT)-$(COMMIT_HASH)"
	zip -j adr-windows-arm64.zip build/windows-arm64/$(BINARY_NAME).exe
	rm -rf build/windows-arm64
dists: binary_linux_amd64 binary_linux_arm64 binary_darwin_amd64 binary_darwin_arm64 binary_windows_amd64 binary_windows_arm64


# Clean build artifacts
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe $(BINARY_NAME)-*

# Clean Go build caches
cleancache: clean
	go clean -cache -testcache -modcache
