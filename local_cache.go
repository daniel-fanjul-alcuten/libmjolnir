package mjölnir

import (
	"bufio"
	"encoding/gob"
	"io"
	"os"
	"time"
)

type lcdata struct {
	Id      string
	Time    time.Time
	Depends []string
}

type localCache struct {
	path string
	data map[string]lcdata
}

func NewLocalCache(path string) *localCache {
	return &localCache{path: path, data: make(map[string]lcdata)}
}

func (cache *localCache) read() error {

	file, err := os.Open(cache.path)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&cache.data)
	if err != nil {
		if err != io.EOF {
			file.Close()
			return err
		}
	}
	return file.Close()
}

func (cache *localCache) write() error {

	file, err := os.Create(cache.path)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	encoder := gob.NewEncoder(writer)
	err = encoder.Encode(cache.data)
	if err != nil {
		file.Close()
		return err
	}

	err = writer.Flush()
	if err != nil {
		file.Close()
		return err
	}
	return file.Close()
}
