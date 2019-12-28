genpb:
	protoc --proto_path=idl \
		-I/usr/local/include \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:./idl \
		--grpc-gateway_out=logtostderr=true:./idl \
		idl/auto_increment.proto
