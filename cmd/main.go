package main

import (
	"log"
	"messageProcessingSystem/internal/process"
	"messageProcessingSystem/storage"
	"messageProcessingSystem/storage/memory"
	"messageProcessingSystem/storage/sqlite"
	"os"
)

const (
	EnvStorageFilePath = "ENV_STORAGE_FILE_PATH"
	EnvStorageType     = "ENV_STORAGE_TYPE"
)

func main() {

	storageFilePath, storageType := GetEnvStorage()

	paymentStorage := UseStorage(storageFilePath, storageType)

	msgProcessor := process.NewMessagesProcessor(paymentStorage)
	for _, messageRaw := range testMessages {
		if err := msgProcessor.PaymentProcessor([]byte(messageRaw)); err != nil {
			log.Print(err)
		}
	}
}

// получение значений из env
func GetEnvStorage() (storageFilePath, storageType string) {
	return os.Getenv(EnvStorageFilePath), os.Getenv(EnvStorageType)
}

// определение используемого хранилища
func UseStorage(storageFilePath, storageType string) storage.Storage {
	var paymentStorage storage.Storage
	switch storageType {
	case "sqlite":
		if storageFilePath == "" {
			log.Fatalf("file path for storage is not found. Use '%s' for set it", EnvStorageFilePath)
		}
		storageLite := sqlite.NewDatabase(storageFilePath)
		if err := storageLite.InitLiteDatabase(); err != nil {
			log.Fatal(err)
		}
		paymentStorage = storageLite
	case "memory":
		paymentStorage = memory.NewDatabase()
	default:
		log.Printf(`storage type is not found. Using default storage in memory. For switch database use '%s'`, EnvStorageType)
		paymentStorage = memory.NewDatabase()
	}
	return paymentStorage
}

var (
	testMessages = []string{
		`{
			"TypeMessage": "created",
			"UidMessage": "1A",
			"AddressFrom": "43245",
			"AddressTo": "4124",
			"Payment": 5000
		}`,
		`{
			"TypeMessage": "processed",
			"UidMessage": "1A"
		}`,
		`{
			"TypeMessage": "canceled",
			"UidMessage": "1A"
		}`,
		`{
			"TypeMessage": "created",
			"UidMessage": "2A",
			"AddressFrom": "43224245",
			"AddressTo": "41123424",
			"Payment": 500000
		}`,
	}
)
