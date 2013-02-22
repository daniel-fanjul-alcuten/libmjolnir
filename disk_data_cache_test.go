package mjölnir

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestDiskDataCache(t *testing.T) {

	path, err := ioutil.TempDir("", "mjolnir-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path)

	cache := NewDiskDataCache(path)

	for _, key := range []string{"", "1", "12", "123", "1234", "12345", "123456", "1234567"} {
		str := []byte("str")

		value, err := cache.Get([]byte(key))
		if err != nil {
			t.Fatal(err)
		}
		if value != nil {
			t.Error("Unexpected Get", value)
		}

		err = cache.Set([]byte(key), str)
		if err != nil {
			t.Fatal(err)
		}

		value, err = cache.Get([]byte(key))
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(value, str) {
			t.Error("Unexpected Get", value)
		}
	}
}
