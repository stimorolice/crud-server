LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/v1
	protoc --proto_path api/v1 \
	--go_out=pkg/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/v1/note.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/server/main.go

send:
	scp service_linux root@89.110.76.162:/root

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t stimorolice/crud-server .
	docker login -u stimorolice -p someshittyasalways##999
	docker image push stimorolice/crud-server 