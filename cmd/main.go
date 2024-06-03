package main

import (
	"fmt"
	messageProcessingSystem "messageProcessingSystem/internal/process"
	memory "messageProcessingSystem/storage"
)

func main() {
	var msg = &memory.Message{}
	if err := messageProcessingSystem.Reader("messages/file1.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Reader("messages/file2.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Reader("messages/file3.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Reader("messages/file4.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}
}
