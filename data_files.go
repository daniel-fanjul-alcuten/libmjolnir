package mjölnir

import (
	"crypto/sha1"
	"encoding/hex"
)

func hashFile(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func getFile(cache dataCache, hash string) ([]byte, error) {
	return cache.Get([]byte(hash))
}

func setFile(cache dataCache, hash string, data []byte) error {
	return cache.Set([]byte(hash), data)
}
