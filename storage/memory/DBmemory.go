package memory

import (
	"fmt"
	"messageProcessingSystem/internal/model"
)

type DBMemory struct {
	inMemory map[string]model.Message
}

func NewDatabase() *DBMemory {
	return &DBMemory{
		inMemory: make(map[string]model.Message),
	}
}

// сохранение данных в базу даннях в памяти
func (db *DBMemory) SavePayment(msg *model.Message) error {

	if _, ok := db.inMemory[msg.UidMessage]; ok {
		return nil
	} else {
		db.inMemory[msg.UidMessage] = *msg
	}

	fmt.Println(db.inMemory)

	return nil
}
