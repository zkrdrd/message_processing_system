package main

import (
	"log"
	"messageProcessingSystem/cmd/app"
	"messageProcessingSystem/internal/process"
)

func main() {

	storageFilePath, storageType := app.GetEnvStorage()

	paymentStorage, err := app.UseStorage(storageFilePath, storageType)
	if err != nil {
		log.Fatal(err)
	}

	msgProcessor := process.NewMessagesProcessor(paymentStorage)
	for _, messageRaw := range testMessages {
		if err := msgProcessor.PaymentProcessor([]byte(messageRaw)); err != nil {
			log.Print(err)
		}
	}
}

var (
	testMessages = []string{
		`{
			"TypeMessage": "created",
			"UidMessage": "1A",
			"AddressFrom": "43245",
			"AddressTo": "4124",
			"Amount": 5000
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
			"Amount": 500000
		}`,
	}
)
