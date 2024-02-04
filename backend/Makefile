.PHONY: gen-proto

# Directory containing .proto files
PROTO_DIR := protos

# Output directory for generated Go files
GO_OUT_DIR := pkg/api

# Find all .proto files in the PROTO_DIR
PROTO_FILES := $(shell find $(PROTO_DIR) -type f -name '*.proto')

# Target for generating Go code from .proto files
gen-proto: $(PROTO_FILES)
	@echo "Generating Go code from .proto files..."
	@mkdir -p $(GO_OUT_DIR)
	@if [ -z "$(PROTO_FILES)" ]; then \
		echo "No .proto files found in $(PROTO_DIR)"; \
		exit 1; \
	fi
	@protoc -I $(PROTO_DIR) \
		--go_out=$(GO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
	@echo "Go code generated successfully!"
