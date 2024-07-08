GO_TOOL=go
BIN_DIR = build
ENTRY_POINT = cmd/main/main.go

all: build run clean
build:
	$(GO_TOOL) build -o $(BIN_DIR)/$(EXEC) $(ENTRY_POINT) 
run: 
	$(GO_TOOL) run $(ENTRY_POINT)
clean: 
	rm -rf $(BIN_DIR)/*

.PHONY: all build run clean