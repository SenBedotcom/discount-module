# Makefile for discount application on Windows

# Variables
APP_NAME = discount.exe
SRC_DIR = .
BUILD_DIR = build

# Go related variables
GO = go
GOFMT = gofmt
GOCMD = $(GO)
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOINSTALL = $(GOCMD) install

# Target: all (default target)
all: build

# Target: build
build:
	@echo "Building the application..."
	$(GOBUILD) -o $(BUILD_DIR)\$(APP_NAME) $(SRC_DIR)

# Target: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Target: fmt
fmt:
	@echo "Formatting Go code..."
	$(GOFMT) -w $(SRC_DIR)

# Target: clean
clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	if exist $(BUILD_DIR)\$(APP_NAME) del $(BUILD_DIR)\$(APP_NAME)
	if exist $(BUILD_DIR) rmdir /s /q $(BUILD_DIR)

# Target: run
run: build
	@echo "Running the application..."
	$(BUILD_DIR)\$(APP_NAME)

# Phony targets
.PHONY: all build test fmt clean run
