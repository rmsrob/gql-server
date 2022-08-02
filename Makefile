BINARY_NAME=gql-server

gql:
	go run github.com/99designs/gqlgen generate

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./client/schema.proto

dep: 
	go mod tidy && go mod vendor && go fmt

run: 
	go run ./server.go

clean: 
	go clean 
	rm ${BINARY_NAME}_linux_amd64