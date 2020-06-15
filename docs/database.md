# BadgerDB

## pkg/nodeconfig

**hostconfig data:**

 This is the configuration for the current styxnode

 ``` Key = pkg/nodeconfig.HostConfigKey (key is a single static value) ```

 ``` Value = pkg/nodeconfig.HostConfigModel ```

**master config data**

This is the master configuration provided to the styxnode

``` Key = "master" ```

``` Value = pkg/nodeconfig.MasterConfigModel ```

**nodeconfig data:**

 nodeconfig is a modified version of the hostconfig.

 nodeconfig is a configuration sent to other styxnodes.

 nodeconfig of other styxnodes are also received and stored in the database.

 ``` Key = "nodeconfig" + "|" + nodeconfig.NodeType + "|" + nodeconfig.NodeID ```

 ``` Value = api/nodeconfig.NodeConfig ```


## pkg/filetransfer