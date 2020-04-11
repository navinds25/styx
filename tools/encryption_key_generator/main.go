// This is for generating a 32 byte key for badgerdb encryption
// refer https://stackoverflow.com/questions/21160258/golang-generating-a-32-byte-key
// refer https://github.com/dgraph-io/badger/pull/1042/files badger/cmd/bank.go

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)
	b64string := base64.StdEncoding.EncodeToString(key)
	fmt.Println("Generated Key:")
	fmt.Println(b64string)
}
