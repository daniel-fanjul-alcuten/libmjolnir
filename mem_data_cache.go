package mjölnir

import (
	"bytes"
)

type memDataCache struct {
	keys   [][]byte
	values [][]byte
}

// stupid and very slow implementation
func NewMemDataCache() *memDataCache {
	return &memDataCache{keys: [][]byte{},
		values: [][]byte{}}
}

func (cache *memDataCache) Get(key []byte) ([]byte, error) {
	for i, k := range cache.keys {
		if bytes.Equal(k, key) {
			return cache.values[i], nil
		}
	}
	return nil, nil
}

func (cache *memDataCache) Set(key, value []byte) error {
	for i, k := range cache.keys {
		if bytes.Equal(k, key) {
			cache.values[i] = value
			return nil
		}
	}
	cache.keys = append(cache.keys, key)
	cache.values = append(cache.values, value)
	return nil
}
