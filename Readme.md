
#STYX:
# Goals of the project
* Peer to peer file transfer application that meets compliance standards
* Peer to peer health checks with leader election
* Single binary for easy distribution
* Allows pulling files from external sftp servers
* Communicate through rpc and not through shared state

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
#### Why have a seperate SFTP Server ? Why not use SSH ?
* The developer of this application knows that software is not perfect and it is better for an application to fail than to lose access to a server.
* Also, this comes in handy when implementing a network architecture with different security zones.

## Features:
### External SFTP Pull
- External SFTP pull main process
  * Read config from file/enter through input
  * send config via grpc client
  * server function add config to database
  * read configs from database
  * process configs from database based on cron/scheduling package
  * pull files from external sftp servers based on config
  * handle custom key exchanges - defer to next version
  * update the database with the pulled files
- Debug external SFTP pull
  * list configs added to database
  * delete configs from database
  * process a specific config from database
  * encrypt/decrypt debugging
- External SFTP post processing
  * extract archives
  * gpg decryption
  * gpg decryption windows
- Security
  * test grpc with certificates - especially on windows
  * add encryption for username/password for external sftp servers
### Internal Transfers
- Configuration
- GRPC Processing
- SFTP Pull/Push
- Post processing

## Progress:
* Read config from file/enter through input
  - [ ] external sftp
  - [ ] internal transfer
* target styxnode based on config
  - [ ] internal transfer
* send config via grpc client
  - [ ] external sftp
  - [ ] internal transfer
* server function add config to database
  - [ ] external sftp
  - [ ] internal transfer
* read configs from database
  - [ ] external sftp
  - [ ] internal transfer
* process configs from database based on cron/scheduling package
  - [ ] external sftp
* process configs based on file system events
  - [ ] internal transfer
* pull files from external sftp servers based on config
  - [ ] external sftp
* handle custom key exchanges - defer to next version
  - [ ] external sftp
* update the database with the pulled files
  - [ ] external sftp
  - [ ] internal transfer
* list configs added to database
  - [ ] external sftp
  - [ ] internal transfer
* delete configs from database
  - [ ] external sftp
  - [ ] internal transfer
* process a specific config from database
  - [ ] external sftp
  - [ ] internal transfer
* extract archives
  - [ ] external sftp
  - [ ] internal transfer
* gpg decryption
  - [ ] external sftp
  - [ ] internal transfer
* gpg decryption windows
  - [ ] external sftp
* test grpc with certificates - especially on windows
  - [ ] external sftp
  - [ ] internal transfer
* add encryption for username/password for external sftp servers
  - [ ] external sftp



## Notes:
* dep ensure/init does not work if the protobuf package is generated. Run make clean before dep ensure it runs to completion.
