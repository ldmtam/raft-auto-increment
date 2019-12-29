BIN = raft-auto-increment

clean:
	rm -f $(BIN)

build: clean
	GOARCH=amd64 go build -o $(BIN)

genpb:
	protoc --proto_path=idl \
		-I/usr/local/include \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:./auto_increment/pb \
		--grpc-gateway_out=logtostderr=true:./auto_increment/pb \
		idl/auto_increment.proto

run-local:
	go run main.go \
		--http-addr localhost:3000 \
		--grpc-addr localhost:4000 \
		--raft-addr localhost:5000 \
		--raft-dir ./mock-raft \
		--data-dir ./mock-data \
		--bootstrap true