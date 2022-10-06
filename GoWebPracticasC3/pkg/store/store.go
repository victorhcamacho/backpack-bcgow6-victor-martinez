package store

import (
	"encoding/json"
	"os"
)

type StoreType string

type fileStore struct {
	FilePath string
}

const (
	FileType StoreType = "file"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

func New(stroe StoreType, path string) Store {
	switch stroe {
	case FileType:
		return &fileStore{path}
	}
	return nil
}

func (fs *fileStore) Write(data interface{}) error {

	if fileData, err := json.MarshalIndent(data, "", "  "); err != nil {
		return err
	} else {
		return os.WriteFile(fs.FilePath, fileData, 0644)
	}
}

func (fs *fileStore) Read(data interface{}) error {

	if file, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	} else {
		return json.Unmarshal(file, &data)
	}

}
