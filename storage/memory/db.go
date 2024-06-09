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

	return nil
}

func (db *DBMemory) GetPaymentById(id string) error {
	for k, m := range db.inMemory {
		if k == id {
			fmt.Println(m)
		}
	}
	return nil
}

func (db *DBMemory) CheckDatabaseAndModelIsCorrect(msg *model.Message) error {
	for k, m := range db.inMemory {
		if k == msg.UidMessage {
			if m.TypeMessage == msg.TypeMessage {
				return fmt.Errorf("model is exist")
			}
		}
		if k != msg.UidMessage {
			if msg.AddressFrom != "" && msg.AddressTo != "" && msg.Payment > 0 {
				return nil
			} else {
				return fmt.Errorf("model is not full")
			}
		}
	}

	return nil
}
