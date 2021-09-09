# Variables
filename ?= foo

# Commands 
run:
	@go run cmd/http/main.go

test:
	@go test ./entity/... ./pkg/... ./repository/... ./service/... ./validate/...

cover:
	@go test ./entity/... ./pkg/... ./repository/... ./service/... ./validate/... -coverprofile coverage.out
	@go tool cover -html=coverage.out

migrate:
	@go run cmd/migrate/main.go

migration:
	@goose -dir=resource/migration create $(filename) sql

gentls:
	@openssl genrsa -out server.key 2048
	@openssl ecparam -genkey -name secp384r1 -out server.key
	@openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 \
		-subj "/C=ID/ST=Jawa Tengah/L=Pemalang/O=Ancene/OU=IT/CN=ancene.org/emailAddress=anceneorg@gmail.com"
	@echo "Success generate 'server.key' and 'server.crt'"

protoc:
	@protoc --go_out=plugins=grpc:proto/. proto/todo.proto

mock:
	@mockery --output=mock --dir=repository --all
