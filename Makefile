

SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_GRPC_OUT := "internal/gRPC"
PKG := "github.com/sim4life/grpc_app_demo"
# PKG := "."
SERVER_PKG_BUILD := "${PKG}/server/gRPC"
CLIENT_PKG_BUILD := "${PKG}/client/gRPC"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all api_grpc build_server build_client

all: build_server build_client

$(API_GRPC_OUT)/:
		mkdir -p $@
								
$(API_GRPC_OUT)/api/grpc/domain.pb.go: api/grpc/domain/*.proto | $(API_GRPC_OUT)/
				@protoc -I api/gRPC \
								-I${GOPATH}/pkg/mod \
								--go_out=plugins=grpc,paths=source_relative:./$(API_GRPC_OUT) \
								api/gRPC/domain/*.proto

$(API_GRPC_OUT)/api/grpc/service.pb.go: api/grpc/service/*.proto | $(API_GRPC_OUT)/
				@protoc -I api/gRPC \
								-I${GOPATH}/pkg/mod \
								--go_out=plugins=grpc,paths=source_relative:./$(API_GRPC_OUT) \
								api/gRPC/service/*.proto

api_grpc: $(API_GRPC_OUT)/api/grpc/domain.pb.go $(API_GRPC_OUT)/api/grpc/service.pb.go ## Auto-generate grpc go sources

dep: ## Get the dependencies
				@go mod download

build_server: dep api_grpc ## Build the binary file for server
				@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build_client: dep api_grpc ## Build the binary file for client
			 	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)
				 
clean: ## Remove previous builds
			 	@rm -rf $(SERVER_OUT) $(CLIENT_OUT) $(API_GRPC_OUT)
				 
help: ## Display this help screen
			 	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
