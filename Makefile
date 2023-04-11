GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
#mkdir ./release

.PHONY: 
all: clean

.PHONY: 
clean:
	rm -rf release
	$(GOCLEAN)

.PHONY: 
ms: 
	go build -v -a -o ./release/bin/monitorService ./cmd/webService/main.go

.PHONY: 
build:
	
.PHONY:
grpc: logGrpc userGrpc

.PHONY:
logGrpc:
	protoc --go_out=./internal/logService/proto/ ./internal/logService/proto/log.proto
	protoc --go-grpc_out=./internal/logService/proto/ ./internal/logService/proto/log.proto
	@echo "genetated log grpc proto successfully!\n"

.PHONY:
userGrpc:
	protoc  --go_out=./internal/userService/proto/ ./internal/userService/proto/user.proto
	protoc  --go-grpc_out=./internal/userService/proto/ ./internal/userService/proto/user.proto
	@echo "genetated user grpc proto successfully!\n"

.PHONY:
prepare:
	echo  "10.182.34.112 etcd.test.com" >> /etc/hosts
	echo  "192.168.0.100 etcd.test.com" >> /etc/hosts
	echo  "192.168.0.100 mysql.test.com">> /etc/hosts
	echo  "10.182.34.112 mysql.test.com">> /etc/hosts
	echo  "10.182.34.112 mysql.test.com">> /etc/hosts
	echo  "192.168.0.100 mongo.test.com">> /etc/hosts


.PHONY:
test:
	go test -v -cover ./...
