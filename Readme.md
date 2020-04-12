
# STYX:
## Goals of the project
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
* This is a cross platform application and will be difficult to handle on windows.

## Features:
### External SFTP Pull
- External SFTP pull main process
  * Read config from file
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
* Read config from file
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

Handling Failures:
Success -> BadgerDB
Failure -> BadgerDB + Log + Alert


## Architecture:

This application transfers files by executing jobs, the jobs defined pre and post transfer logic for determining how and whether the transfer should take place.
Each job consists of Actions in order to do so. These are the following actions 
1. Trigger: This handles the start of the job.
  - [ ] Manual Request
  - [ ] Cron
  - [ ] Inotify
2. PreExecutor:
  This is for executing something before the start of a job. eg: send http request or whitelist connection or run a check.
  - [ ] Transfer Check - checks that it's okay to transfer from the directory defined in the job.
3. Matcher: This matches files based on timestamp and/or pattern. Need to change this from switch-case to if-else.
  - [ ] Transfer Check - checks that it's okay to transfer from the directory defined in the job.
  - [ ] Matcher
4. FileCheck: This runs a check on individual files that are about to be transferred.
  - [ ] OL - OL or Overwrite Logic handles the logic for files that are about to be overwritten.
5. Transfer: 
  - [ ] Pull
  - [ ] Push
  - [ ] PullExternal
  - [ ] PushExternal
6. PostExecutor: This is for Executing something at the end of a job.
  - [ ] CmdExecutor - for executing a command.


### Job Configuration:

* Trigger
* Source
- NodeID -
- SPath
- SFile
- SCondition
- SExecutor
* Destination - 

### Node/Peer Configuration:

* NodeID
- External - bool
- IPAddress/DomainName
- GRPC Port
- SFTP Port
- GRPC TLS Key
- GRPC TLS Cert
- GRPC Authentication ?
- SFTP Username
- SFTP Password


### Flow:

* Setup:
1. logging
2. on styxnode, Cli -> NodeConfigSelf
3. on styxmaster, add nodecofig of new node.
4. styxmaster sends NodeConfig to styxnode

* FTJobs:
1. create new job config - (need to address this)
2. add FTJob from styxmaster
3. Determines add/run based on Trigger.
 - Runjob: Determines job type, runs on appropriate styxnode
 - Addjob: Determines job type, adds on appropriate styxnode

* RunJob flow: