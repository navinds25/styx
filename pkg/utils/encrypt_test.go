package utils

import (
	"testing"
)

// TestAESEncryptCBC is for testing the AESEncryptCBC function.
func TestAESEncryptCBC(t *testing.T) {
	key := "exbuVWWdo29WD3Ri"
	msg := "secret phrase"
	ciphertxt, err := AESEncryptCBC([]byte(key), []byte(msg))
	if err != nil {
		t.Error(err)
	}
	plaintxt, err := AESDecryptCBC([]byte(key), ciphertxt)
	if err != nil {
		t.Error(err)
	}
	if string(plaintxt) != msg {
		t.Fatalf("got: %s instead of %s", string(plaintxt), msg)
	}
}
