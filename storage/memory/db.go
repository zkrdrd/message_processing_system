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

// TODO:
// 1. статус платежа Бд не обновляется //
// 2. не проверяется если платеж уже прошел //

// сохранение данных в базу даннях в памяти
func (db *DBMemory) SavePayment(msg *model.Message) error {

	if ok := db.findPaymentForValidate(msg); !ok {
		if err := msg.ValidatePaymenIfNotExistInDB(); err != nil {
			return err
		}
	}

	if val, ok := db.inMemory[msg.UidMessage]; ok {
		if val.TypeMessage != msg.TypeMessage && val.TypeMessage == "created" {
			val.TypeMessage = msg.TypeMessage
		}
		db.inMemory[msg.UidMessage] = val
		return nil
	} else {
		db.inMemory[msg.UidMessage] = *msg
	}

	return nil
}

func (db *DBMemory) GetPaymentById(id string) (*model.GetedPayment, error) {
	if val, ok := db.inMemory[id]; ok {
		return &model.GetedPayment{
			TypeMessage: val.TypeMessage,
			UidMessage:  val.UidMessage,
			AddressFrom: val.AddressFrom,
			AddressTo:   val.AddressTo,
			Payment:     val.Payment,
		}, nil
	}
	return nil, fmt.Errorf(`id is not found`)
}

func (db *DBMemory) findPaymentForValidate(msg *model.Message) bool {
	if _, ok := db.inMemory[msg.UidMessage]; ok {
		return true
	}

	return false
}
