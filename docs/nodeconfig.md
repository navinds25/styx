NodeConfig

github.com/navinds25/styx/pkg/nodeconfig

In order to implement the peer to peer transfer each styxnode has to know the details of the other nodes in order to communicate with them, this information is handled by the styxconfig package.

HostConfig holds the configuration of the styxnode when it starts, this includes data the other styxnodes do not require for peer to peer communication.
This is saved in the config database as hostconfig:key_name and value.

NodeConfig handles the configuration of the styxnode for the peer to peer communication.

styxnode -config hostconfig 1.req [ (self) nodeconfig -grpc-> styxnode ] response [ all nodeconfig <- styxnode]
----
server a

