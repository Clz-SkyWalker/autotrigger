cd ./log/rpc/
goctl rpc protoc ./log.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
::protoc ./log.proto --go_out=./types --go-grpc_out=./types
cd ../../../
