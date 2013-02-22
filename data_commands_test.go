package mjölnir

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHashCommand(t *testing.T) {

	time := "time"
	exec := "example"
	args := [][]string{
		[]string{"-a", "b"},
		[]string{"-a", "cache"},
		[]string{"-z"}}
	inputs := []string{"a", "z"}
	example := &commandKey{time, exec, args, inputs}

	h := sha1.New()
	writeString(h, time)
	writeString(h, exec)
	writeStringSliceSlice(h, args)
	writeStringSlice(h, inputs)

	hash := hashCommand(example)
	if hash != hex.EncodeToString(h.Sum(nil)) {
		t.Error("Unexpected hash", hash)
	}
}

func TestSetCommand(t *testing.T) {

	key := &commandKey{}
	cache := NewMemDataCache()

	value, err := getCommand(cache, key)
	if err != nil {
		t.Fatal(err)
	}
	if value != nil {
		t.Error("Unexpected value", value)
	}

	value = &commandValue{[]byte("foo"), []byte("bar"), []string{"f1", "f2"}, 2}
	err = setCommand(cache, key, value)
	if err != nil {
		t.Fatal(err)
	}

	value, err = getCommand(cache, key)
	if err != nil {
		t.Fatal(err)
	}
	if value == nil {
		t.Fatal("Unexpected value", value)
	}
	if !bytes.Equal(value.Stdout, []byte("foo")) {
		t.Error("Unexpected Stdout", value.Stdout)
	}
	if !bytes.Equal(value.Stderr, []byte("bar")) {
		t.Error("Unexpected Stderr", value.Stderr)
	}
	if fmt.Sprint(value.Files) != fmt.Sprint([]string{"f1", "f2"}) {
		t.Error("Unexpected Files", value.Files)
	}
	if value.Status != 2 {
		t.Error("Unexpected Status", value.Status)
	}
}
