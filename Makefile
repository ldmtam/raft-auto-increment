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
		--id 1 \
		--addr localhost:3000 \
		--raft-addr localhost:5000 \
		--raft-dir ./tmp/node1/raft \
		--data-dir ./tmp/node1/data \
		--bootstrap true

run2:
	go run main.go \
		--id 2 \
		--addr localhost:13000 \
		--raft-addr localhost:15000 \
		--join-addr localhost:3000 \
		--raft-dir ./tmp/node2/raft \
		--data-dir ./tmp/node2/data \

run3:
	go run main.go \
		--id 3 \
		--addr localhost:23000 \
		--raft-addr localhost:25000 \
		--join-addr localhost:3000 \
		--raft-dir ./tmp/node3/raft \
		--data-dir ./tmp/node3/data \
