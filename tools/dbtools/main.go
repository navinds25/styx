package main

import (
	"bytes"
	"encoding/gob"
	"strings"

	log "github.com/sirupsen/logrus"

	badger "github.com/dgraph-io/badger/v2"
)

// BadgerDB is the DB instance for BadgerDB
type BadgerDB struct {
	DB *badger.DB
}

// DataStore is the struct containing the NodeConfigStore interface
type DataStore struct {
	ToolDB Store
}

// Data is the instance of DataStore
var Data DataStore

// Store is the interface for all NodeConfig DB Actions
type Store interface {
	AddEntry() error
	GetEntry() error
	CloseDB() error
}

// InitDB initializes the NodeConfigDB
func InitDB(s Store) {
	Data.ToolDB = s
}

type bigVal struct {
	Key1 string
	Key2 int32
	Key3 string
}

func (bd BadgerDB) AddEntry() error {
	key := strings.TrimSpace("hello")
	value := &bigVal{
		Key1: "something", Key2: 32,
	}
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		log.Error(err)
	}
	txn := bd.DB.NewTransaction(true)
	defer txn.Discard()
	if err := txn.Set([]byte(key), buf.Bytes()); err != nil {
		log.Fatal(err)
	}
	if err := txn.Commit(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (bd BadgerDB) GetEntry() error {
	txn := bd.DB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte("hello"))
	if err != nil {
		log.Error(err)
	}
	val, err := item.ValueCopy(nil)
	if err != nil {
		log.Error(err)
	}
	out := &bigVal{}
	if err := gob.NewDecoder(bytes.NewReader(val)).Decode(out); err != nil {
		log.Error(err)
	}
	log.Println(out.Key1)
	return nil
}

func (bd BadgerDB) CloseDB() error {
	return nil
}

func startDB() {
	dbOpts := badger.DefaultOptions("/home/navin/Coding/go/styx/tools/dbtools/data")
	dbInst, err := badger.Open(dbOpts)
	if err != nil {
		log.Fatal(err)
	}
	InitDB(BadgerDB{
		DB: dbInst,
	})
}

func main() {
	startDB()
	err := Data.ToolDB.AddEntry()
	if err != nil {
		log.Fatal(err)
	}
	err = Data.ToolDB.GetEntry()
	if err != nil {
		log.Fatal(err)
	}
}
