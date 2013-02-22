package mjölnir

import (
	"bytes"
	"testing"
)

func TestMemDataCache(t *testing.T) {

	cache := NewMemDataCache()

	foo := []byte("foo")
	bar := []byte("bar")

	value, err := cache.Get(foo)
	if err != nil {
		t.Fatal(err)
	}
	if value != nil {
		t.Error("Unexpected Get", value)
	}

	err = cache.Set(foo, bar)
	if err != nil {
		t.Fatal(err)
	}

	value, err = cache.Get(foo)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(value, bar) {
		t.Error("Unexpected Get", value)
	}
}
