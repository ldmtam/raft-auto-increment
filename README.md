# Raft Auto Increment
Distributed, fault-tolerant, persistent auto increment service with Raft consensus. <br/>

Support both REST API and gRPC.

# What is Raft consensus?
Please refer to this [page](https://raft.github.io/) for more detail.

# How to run
Clone the project:
```
go get github.com/ldmtam/raft-auto-increment
```

Change directory to `raft-auto-increment`, create `tmp` folder and `node1`, `node2`, `node3` folder inside `tmp`.

Start node 1, this node will be the leader by default.
```
make run1
```

Start other 2 nodes by running following commands, these 2 nodes will be slaves to node 1
```
make run2
make run3
```

Get next available ID for key `foo`:
```
curl http://localhost:3000/auto-increment/one/foo
```

Get last inserted ID for key `foo`:
```
curl http://localhost:3000/auto-increment/last-inserted/foo
```

Get next 10 available IDs for key `bar`:
```
curl http://localhost:3000/auto-increment/many/bar/10
```
