run:
	go run cmd/rusprofile/main.go

gen:
	protoc -I=internal/proto \
	--go_out=internal/proto \
	--go-grpc_out=internal/proto \
	--grpc-gateway_out=logtostderr=true:internal/proto  \
	--openapiv2_out=logtostderr=true:internal/proto/swagger \
	rusprofile.proto

lint:
	clang-format -n --style=Google internal/proto/*.proto
