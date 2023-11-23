# golang-grpc

Códigos do curso da udemy sobre gRPC em golang

protoc -Igreet/proto --go_opt=module=golang-grpc --go_out=. --go-grpc_opt=module=golang-grpc  --go-grpc_out=. ./greet/proto/dummy.proto
