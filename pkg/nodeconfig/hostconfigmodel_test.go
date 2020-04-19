package nodeconfig_test

import (
	"reflect"
	"testing"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/navinds25/styx/pkg/nodeconfig"
)

func TestBadgerDB_AddGetHostConfigEntry(t *testing.T) {
	// initialize db
	opts := badger.DefaultOptions("").WithInMemory(true)
	ncdb, err := badger.Open(opts)
	if err != nil {
		t.Error("error creating db", err)
	}
	nodeconfig.InitNodeConfigDB(nodeconfig.BadgerDB{NodeConfigDB: ncdb})
	key := "test1"
	value := &nodeconfig.HostConfigModel{
		NodeID:    "test1",
		NodeType:  "internal",
		IPAddress: "127.0.0.1",
		GRPCPort:  28888,
		SFTPPort:  28889,
		SZ:        "app",
		SFTPAuth: nodeconfig.SFTPAuthModel{
			SFTPAuthType: "",
			HostkeyFile:  "ssh_host_rsa_key",
		},
		ExternalAccess: true,
	}
	if err := nodeconfig.Data.NodeConfig.AddHostConfigEntry(key, value); err != nil {
		t.Error("error adding data to db")
	}
	hcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(key)
	if err != nil {
		t.Error("error getting value from db")
	}
	if !reflect.DeepEqual(hcM, value) {
		t.FailNow()
	}
}

//tests := []struct {
//	name    string
//	fields  fields
//	args    args
//	wantErr bool
//}{
//	// TODO: Add test cases.
//}
//for _, tt := range tests {
//	t.Run(tt.name, func(t *testing.T) {
//		badgerDB := BadgerDB{
//			NodeConfigDB: tt.fields.NodeConfigDB,
//		}
//		if err := badgerDB.AddHostConfigEntry(tt.args.id, tt.args.inModel); (err != nil) != tt.wantErr {
//			t.Errorf("BadgerDB.AddHostConfigEntry() error = %v, wantErr %v", err, tt.wantErr)
//		}
//	})
//}
