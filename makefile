# Makefile

# Define the Go source file
GO_SOURCE := main.go

# Define the output binary name
OUTPUT_BINARY := rp_mgmt

# Define the output binary name
APPNAME := rp_mgmt

# Default target
all: build test

# Target to build the Go project
build: prepare
	go build -o "./$(OUTPUT_BINARY)" $(GO_SOURCE)

run:
	go build -o "./$(OUTPUT_BINARY)" $(GO_SOURCE)
	./$(OUTPUT_BINARY)

test:
	go test ./... -v

init:
	go mod int $(APPNAME)

prepare:
	go mod tidy