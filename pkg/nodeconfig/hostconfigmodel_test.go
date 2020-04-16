package nodeconfig

import (
	"reflect"
	"testing"

	badger "github.com/dgraph-io/badger/v2"
)

func TestBadgerDB_AddGetHostConfigEntry(t *testing.T) {
	// initialize db
	opts := badger.DefaultOptions("").WithInMemory(true)
	ncdb, err := badger.Open(opts)
	if err != nil {
		t.Error("error creating db", err)
	}
	InitNodeConfigDB(BadgerDB{NodeConfigDB: ncdb})
	key := "test1"
	value := &HostConfigModel{
		NodeID:    "test1",
		NodeType:  "internal",
		IPAddress: "127.0.0.1",
		GRPCPort:  28888,
		SFTPPort:  28889,
		SZ:        "app",
		SFTPAuth: SFTPAuthModel{
			SFTPAuthType: "",
			HostkeyFile:  "ssh_host_rsa_key",
		},
		ExternalAccess: true,
	}
	if err := Data.NodeConfig.AddHostConfigEntry(key, value); err != nil {
		t.Error("error adding data to db")
	}
	hcM, err := Data.NodeConfig.GetHostConfigEntry(key)
	if err != nil {
		t.Error("error getting value from db")
	}
	if reflect.DeepEqual(hcM, value) {
		t.Log("worked")
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
