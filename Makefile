.PHONY: run
run:
	go run cmd/rusprofile/main.go

.PHONY: client
client:
	go run cmd/client/client.go

.PHONY: gen
gen:
	protoc -I=api/proto \
	--go_out=api/gen \
	--go-grpc_out=api/gen \
	--grpc-gateway_out=logtostderr=true:api/gen  \
	--openapiv2_out=logtostderr=true:api/gen \
	rusprofile.proto

lint:
	clang-format -n --style=Google internal/proto/*.proto
