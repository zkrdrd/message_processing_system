package memory

import (
	"fmt"
	message "messageProcessingSystem/storage"
)

type DBMemory struct {
	inMemory map[string]message.Message
}

func NewDatabase() *DBMemory {
	return &DBMemory{
		inMemory: make(map[string]message.Message),
	}
}

// сохранение данных в базу даннях в памяти
func (db *DBMemory) SavePayment(msg *message.Message) error {

	if _, ok := db.inMemory[msg.UidMessage]; ok {
		return nil
	} else {
		db.inMemory[msg.UidMessage] = *msg
	}

	fmt.Println(db.inMemory)

	return nil
}
