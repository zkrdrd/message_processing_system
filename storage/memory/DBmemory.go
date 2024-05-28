package memory

import (
	"fmt"
	message "messageProcessingSystem/storage"
)

var inMemory = make(map[string]message.Message)

// сохранение данных в базу даннях в памяти
func SavePayment(msg *message.Message) error {

	inMemory[msg.UidMessage] = *msg
	fmt.Println(inMemory)

	return nil
}
