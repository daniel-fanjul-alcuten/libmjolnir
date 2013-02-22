package mjölnir

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

type diskDataCache struct {
	path string
}

func NewDiskDataCache(path string) *diskDataCache {
	return &diskDataCache{path}
}

func (cache *diskDataCache) Get(key []byte) ([]byte, error) {

	dirname, filename := cache.getFilePath(cache.path, key)

	data, err := ioutil.ReadFile(filepath.Join(dirname, filename))
	if os.IsNotExist(err) {
		return nil, nil
	}
	return data, err
}

func (cache *diskDataCache) Set(key, value []byte) error {

	dirname, filename := cache.getFilePath(cache.path, key)

	err := os.MkdirAll(dirname, 0777)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(dirname, filename), value, 0666)
}

func (cache *diskDataCache) getFilePath(path string, key []byte) (string, string) {

	str := hex.EncodeToString(key)

	if len(str) > 6 {
		dirname := str[:2] + string(os.PathSeparator) + str[2:4] + string(os.PathSeparator) + str[4:6]
		return filepath.Join(path, dirname), str[6:]

	} else if len(str) > 5 {
		dirname := str[:2] + string(os.PathSeparator) + str[2:4] + string(os.PathSeparator) + str[4:6]
		return filepath.Join(path, dirname), "_"

	} else if len(str) > 4 {
		dirname := str[:2] + string(os.PathSeparator) + str[2:4]
		return filepath.Join(path, dirname), str[4:]

	} else if len(str) > 3 {
		dirname := str[:2] + string(os.PathSeparator) + str[2:4]
		return filepath.Join(path, dirname), "_"

	} else if len(str) > 2 {
		dirname := str[:2]
		return filepath.Join(path, dirname), str[2:]

	} else if len(str) > 1 {
		dirname := str[:2]
		return filepath.Join(path, dirname), "_"
	}

	return path, "_"
}
