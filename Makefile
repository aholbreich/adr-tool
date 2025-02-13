BINARY_NAME=adr
INSTALL_DIR=$(HOME)/bin

build:
	go build -o $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_DIR)
	mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
