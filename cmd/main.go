package main

import (
	"fmt"
	"messageProcessingSystem/internal/model"
	messageProcessingSystem "messageProcessingSystem/internal/process"

	"github.com/zkrdrd/ConfigParser"
)

func main() {
	var msg = &model.Message{}

	if err := ConfigParser.Read("messages/file1.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}

	if err := ConfigParser.Read("messages/file2.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}

	if err := ConfigParser.Read("messages/file3.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}

	if err := ConfigParser.Read("messages/file4.json", msg); err != nil {
		fmt.Println(err)
	}
	if err := messageProcessingSystem.Processing(msg); err != nil {
		fmt.Println(err)
	}
}
