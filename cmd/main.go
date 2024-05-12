package main

import (
	"fmt"
	"messageProcessingSystem"
)

/*type Message struct {
	TypeMessage string `json:"TypeMesage"`
	UidMessage  string `json:"UidMessage"`
	AddressFrom string `json:"AddressFrom,omitempty"`
	AddressTo   string `json:"AddressTo,omitempty"`
	Payment     int    `json:"Payment,omitempty"`
}*/

func main() {
	//var msg = &Message{}
	var msg map[string]interface{}
	if err := messageProcessingSystem.Processing("messages/path1.json", &msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}
