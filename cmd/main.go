package main

import (
	"fmt"
	messageProcessingSystem "messageProcessingSystem/internal/process"
	storage "messageProcessingSystem/storage/lite"
	memory "messageProcessingSystem/storage/memory"
)

func main() {
	var msg = &memory.Message{}
	//var msg map[string]interface{}
	storage.NewSqlite()
	if err := messageProcessingSystem.Processing("messages/path1.json", &msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}
