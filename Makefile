BIN = raft-auto-increment

clean:
	rm -f $(BIN)

build: clean
	CGO_ENABLE=0 go build -o $(BIN)

genpb:
	buf generate

run1:
	go run main.go \
		--raft-id 1 \
		--node-addr localhost:3000 \
		--raft-addr localhost:5000 \
		--raft-dir ./tmp/node1/raft \
		--storage bitcask \
		--bootstrap true

run2:
	go run main.go \
		--raft-id 2 \
		--node-addr localhost:13000 \
		--raft-addr localhost:15000 \
		--join-addr localhost:3000 \
		--raft-dir ./tmp/node2/raft \
		--storage badger

run3:
	go run main.go \
		--raft-id 3 \
		--node-addr localhost:23000 \
		--raft-addr localhost:25000 \
		--join-addr localhost:3000 \
		--raft-dir ./tmp/node3/raft \
		--storage bitcask
