package storage

import "fmt"

type StorageFactory struct{}

func NewStorageFactory() *StorageFactory {
	return &StorageFactory{}
}

func (f *StorageFactory) CreateStorage(storageType string, filePath string) (Operator, error) {
	switch storageType {
	case "json":
		return NewJsonStorage(filePath), nil
	case "yaml":
		return NewYamlStorage(filePath), nil
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", storageType)
	}
}
