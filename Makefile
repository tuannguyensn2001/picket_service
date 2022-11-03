install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    	google.golang.org/protobuf/cmd/protoc-gen-go \
    	google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen-proto:
	@mkdir -p src/pb/${name}
	@rm -f src/pb/${name}/*.go
	@protoc --proto_path=proto --go_out=src/pb/${name} --go_opt=paths=source_relative \
	--go-grpc_out=src/pb/${name} --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=src/pb/${name} --grpc-gateway_opt=paths=source_relative \
	proto/${name}.proto

build-err:
	@go run src/server/main.go build-err

gen-err:
	@go run src/server/main.go gen-error