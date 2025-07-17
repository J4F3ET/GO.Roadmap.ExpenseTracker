BINARY_NAME=expense-tracker

# This is the directory where the binary will be placed
# after building the project.
OUTPUT_DIR=bin

# Entry point for the Go application.
ENTRY=main.go

# This is the target to clean the build directory.
build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(ENTRY)

run: build
	./$(OUTPUT_DIR)/$(BINARY_NAME)

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin/