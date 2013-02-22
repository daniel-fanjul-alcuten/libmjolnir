package mjölnir

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"io"
	"sort"
)

type commandKey struct {
	id     string
	exec   string
	args   [][]string
	inputs []string
}

type commandValue struct {
	Stdout []byte
	Stderr []byte
	Files  []string
	Status int
}

func writeString(w io.Writer, value string) {
	binary.Write(w, binary.BigEndian, len(value))
	w.Write([]byte(value))
}

func writeStringSlice(w io.Writer, values []string) {
	binary.Write(w, binary.BigEndian, len(values))
	for _, value := range values {
		writeString(w, value)
	}
}

func writeStringSliceSlice(w io.Writer, values [][]string) {
	binary.Write(w, binary.BigEndian, len(values))
	for _, value := range values {
		writeStringSlice(w, value)
	}
}

type stringSliceSlice [][]string

func (p stringSliceSlice) Len() int {
	return len(p)
}

func (p stringSliceSlice) Less(i, j int) bool {
	index := 0
	for index < len(p[i]) && index < len(p[j]) {
		if p[i][index] < p[j][index] {
			return true
		} else if p[i][index] > p[j][index] {
			return false
		}
		index++
	}
	return len(p[i]) < len(p[j])
}

func (p stringSliceSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func hashCommand(key *commandKey) string {

	id := key.id
	exec := key.exec
	args := key.args
	inputs := key.inputs

	sort.Sort(stringSliceSlice(args))
	sort.Strings(inputs)

	hasher := sha1.New()
	writeString(hasher, id)
	writeString(hasher, exec)
	writeStringSliceSlice(hasher, args)
	writeStringSlice(hasher, inputs)

	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func getCommand(cache dataCache, key *commandKey) (*commandValue, error) {

	hash := hashCommand(key)

	data, err := cache.Get([]byte(hash))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	var value commandValue
	return &value, gob.NewDecoder(bytes.NewBuffer(data)).Decode(&value)
}

func setCommand(cache dataCache, key *commandKey, value *commandValue) error {

	hash := hashCommand(key)

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	if err != nil {
		return err
	}

	return cache.Set([]byte(hash), buf.Bytes())
}
