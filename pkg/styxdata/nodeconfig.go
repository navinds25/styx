package styxdata

type StyxNode struct {
	NodeID    string
	IPAddress string
	GRPCPort  string
	SFTPPort  string
	SZ        int    // Security Zone 0 = dmz 3 = data segment
	HostKey   []byte // Need to test this.

}

// AddNodeEntry Adds an entry for a new styx node
func (badgerDB BadgerDB) AddNodeEntry() error {
	return nil
}
