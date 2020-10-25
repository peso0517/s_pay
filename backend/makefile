GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
# GOTEST=$(GOCMD) test
BINARY_NAME=binary

all: clean run
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
# test:
#     $(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
