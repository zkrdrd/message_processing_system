package main

import (
	"fmt"
	messageProcessingSystem "messageProcessingSystem/internal/process"
	memory "messageProcessingSystem/storage"
	storage "messageProcessingSystem/storage/lite"
)

func main() {
	var msg = &memory.Message{}
	storage.InitDatabase()
	if err := messageProcessingSystem.Reader("messages/file1.json", msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	messageProcessingSystem.Processing(msg)
}
