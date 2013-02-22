package utils

import (
	"io/ioutil"
	"os"
)

type TempDir struct {
	names map[string]bool
}

func NewTempDir() *TempDir {
	return &TempDir{make(map[string]bool)}
}

func (t *TempDir) newFile() (*os.File, error) {
	file, err := ioutil.TempFile("", "mjolnir-")
	if err != nil {
		return nil, err
	}
	name := file.Name()
	t.names[name] = true
	return file, nil
}

func (t *TempDir) NewFileData(data []byte) (string, error) {
	file, err := t.newFile()
	if err != nil {
		return "", err
	}
	_, err = file.Write(data)
	if err != nil {
		return "", err
	}
	name := file.Name()
	t.names[name] = true
	return name, file.Close()
}

func (t *TempDir) NewFileName() (string, error) {
	file, err := t.newFile()
	if err != nil {
		return "", err
	}
	name := file.Name()
	t.names[name] = true
	return name, file.Close()
}

func (t *TempDir) Remove() []error {
	errs := []error{}
	for name := range t.names {
		if err := os.Remove(name); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
