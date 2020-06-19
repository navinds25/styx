NodeConfig

github.com/navinds25/styx/pkg/nodeconfig

In order to implement the peer to peer transfer each styxnode has to know the details of the other nodes in order to communicate with them, this information is handled by the styxconfig package.

HostConfig holds the configuration of the styxnode when it starts, this includes data the other styxnodes do not require for peer to peer communication.
This is saved in the config database as hostconfig:key_name and value.

NodeConfig handles the configuration of the styxnode for the peer to peer communication.

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
