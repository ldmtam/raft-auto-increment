# Raft Auto Increment
Distributed, fault-tolerant, persistent, auto-increment ID generation service with Raft consensus.

Support both `REST API` and `gRPC` protocol.

Support different storage databases for Raft log: `Bitcask` and `Badger`

# What is Raft consensus?
Please refer to this [page](https://raft.github.io/) for more detail.

# How to run
Clone the project:
```
git clone https://github.com/ldmtam/raft-auto-increment
```

Change directory to `raft-auto-increment`
```
cd raft-auto-increment
```

Start `node 1`, this node will be the leader by default. `node 1` will serve requests at port `3000`.
```
make run1
```

Start other 2 nodes by running following commands. `node 2` and `node 3` will be slaves of `node 1` and serve requests at port `13000`, `23000` respectively.
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

You can send request to `node 2` and `node 3` as well. Requests will be automatically forwarded to master node (`node 1`),
```
curl http://localhost:13000/auto-increment/one/bar
```

```
curl http://localhost:23000/auto-increment/last-inserted/bar
```
