package main

import (
	"fmt"
	messageProcessingSystem "messageProcessingSystem/internal/process"
	memory "messageProcessingSystem/storage"
	storage "messageProcessingSystem/storage/lite"
)

func main() {
	var msg = &memory.Message{}
	//var msg map[string]interface{}
	//var msg map[string]memory.Message
	storage.NewSqlite()
	if err := messageProcessingSystem.Reader("messages/path1.json", msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	var mmr memory.MessageReader = msg
	a := mmr.GetUid()
	fmt.Println(a)
}
