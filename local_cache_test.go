package mjölnir

import (
	"fmt"
	. "github.com/daniel-fanjul-alcuten/libmjolnir/utils"
	"testing"
	"time"
)

func TestLocalCache(t *testing.T) {

	tempDir := NewTempDir()
	defer tempDir.Remove()

	path, err := tempDir.NewFileName()
	if err != nil {
		t.Fatal(err)
	}

	cache := NewLocalCache(path)
	now := time.Now()
	cache.data["foo"] = lcdata{"bar", now, []string{"baz"}}
	err = cache.write()
	if err != nil {
		t.Fatal(err)
	}

	cache = NewLocalCache(path)
	err = cache.read()
	if err != nil {
		t.Fatal(err)
	}
	if cache.data["foo"].Id != "bar" {
		t.Error("unexpected Time", cache.data["foo"].Time)
	}
	if cache.data["foo"].Time != now {
		t.Error("unexpected Time", cache.data["foo"].Time)
	}
	if fmt.Sprint(cache.data["foo"].Depends) != fmt.Sprint([]string{"baz"}) {
		t.Error("unexpected Depends", cache.data["foo"].Depends)
	}
}
