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

run1:
	go run main.go \
		--raft-id 1 \
		--node-addr localhost:3000 \
		--raft-addr localhost:5000 \
		--raft-dir ./tmp/node1/raft \
		--storage badgerdb \
		--bootstrap true

run2:
	go run main.go \
		--raft-id 2 \
		--node-addr localhost:13000 \
		--raft-addr localhost:15000 \
		--join-addr localhost:3000 \
		--raft-dir ./tmp/node2/raft \
		--storage badgerdb

run3:
	go run main.go \
		--raft-id 3 \
		--node-addr localhost:23000 \
		--raft-addr localhost:25000 \
		--join-addr localhost:3000 \
		--raft-dir ./tmp/node3/raft \
		--storage badgerdb
