package app

import (
	"log"
	"messageProcessingSystem/storage"
	"messageProcessingSystem/storage/memory"
	"messageProcessingSystem/storage/sqlite"
	"os"
)

const (
	EnvStorageFilePath = "ENV_STORAGE_FILE_PATH"
	EnvStorageType     = "ENV_STORAGE_TYPE"
)

// получение значений из env
func GetEnvStorage() (storageFilePath, storageType string) {
	return os.Getenv(EnvStorageFilePath), os.Getenv(EnvStorageType)
}

// определение используемого хранилища
func UseStorage(storageFilePath, storageType string) (storage.Storage, error) {
	var paymentStorage storage.Storage
	switch storageType {
	case "sqlite":
		if storageFilePath == "" {
			log.Fatalf("file path for storage is not found. Use '%s' for set it", EnvStorageFilePath)
		}
		storageLite := sqlite.NewDatabase(storageFilePath)
		if err := storageLite.InitLiteDatabase(); err != nil {
			return nil, err
		}
		paymentStorage = storageLite
	case "memory":
		paymentStorage = memory.NewDatabase()
	default:
		log.Printf(`storage type is not found. Using default storage in memory. For switch database use '%s'`, EnvStorageType)
		paymentStorage = memory.NewDatabase()
	}
	return paymentStorage, nil
}
