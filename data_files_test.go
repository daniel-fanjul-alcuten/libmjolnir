package mjölnir

import (
	"bytes"
	"testing"
)

func TestHashFile(t *testing.T) {

	data := []byte("foo")
	hash := hashFile(data)

	if hash != "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33" {
		t.Error("Unexpected hash", hash)
	}
}

func TestSetFile(t *testing.T) {

	cache := NewMemDataCache()

	data := []byte("foo")
	hash := "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"

	value, err := getFile(cache, hash)
	if err != nil {
		t.Fatal(err)
	}
	if value != nil {
		t.Error("Unexpected value", value)
	}

	err = setFile(cache, hash, data)
	if err != nil {
		t.Fatal(err)
	}

	value, err = getFile(cache, hash)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(value, data) {
		t.Error("Unexpected value", value)
	}
}
