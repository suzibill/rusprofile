run:
	go run cmd/rusprofile/main.go

gen:
	protoc -I=internal/proto --go_out=internal/proto --go-grpc_out=internal/proto rusprofile.proto

lint:
	clang-format -n --style=Google internal/proto/*.proto
