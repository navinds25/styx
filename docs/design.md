## Goals of the project
* Peer to peer file transfer application that meets compliance standards
* Single binary for easy distribution
* Allows pulling files from external sftp servers

## Design decisions:
### Golang:
* Readability
* Easy distribution
* Good support for RPC protocols
* Quick Development speed
* Low level language
### GRPC:
* Bidirectional streaming, allows for client to wait on connections
* Readability, protobuf file is easy to read and get an overall picture
* Binary transmission reduces network load
* Expecting support for GRPC transcoding to improve in future
* Expecting support for QUIC in future
* Why not Websockets with HTTP? - Code becomes overly complicated.
* Why not message pack or simple json ? - GRPC is more readable, better supported in Golang and is faster over the network.
### BadgerDB:
* Golang Native
* Performant
* Why not SQLite ? - SQLite requires cgo, it's not a performant and Styx does not have a requirement that needs sql specific operations like joins.
* Why not boltdb ? - BoltDB is a binary tree database and takes a lot more memory than an LSM database like BadgerDB.
### SFTP Server:
#### Why have a separate SFTP Server ? Why not use SSH ?
* This is a cross platform application and will be difficult to handle on windows.
