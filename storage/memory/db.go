package memory

import (
	"fmt"
	"messageProcessingSystem/internal/model"
)

type DBMemory struct {
	inMemory map[string]model.Message
}

type GetMessage struct {
	TypeMessage string
	UidMessage  string
	AddressFrom string
	AddressTo   string
	Payment     int
	//CreatedAt   string
	//ModifyAt    string
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

	if err := db.checkPaymentIsExist(msg); err != nil {
		return err
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

func (db *DBMemory) GetPaymentById(id string) error {
	if val, ok := db.inMemory[id]; ok {
		gm := &GetMessage{
			TypeMessage: val.TypeMessage,
			UidMessage:  val.UidMessage,
			AddressFrom: val.AddressFrom,
			AddressTo:   val.AddressTo,
			Payment:     val.Payment,
		}
		fmt.Println(gm)
	}
	return nil
}

func (db *DBMemory) checkPaymentIsExist(msg *model.Message) error {
	for k, m := range db.inMemory {
		if k == msg.UidMessage {
			if m.TypeMessage == msg.TypeMessage {
				return fmt.Errorf("model is exist")
			}
		}
		if k != msg.UidMessage && (msg.AddressFrom == "" || msg.AddressTo == "" || msg.Payment <= 0) {
			return fmt.Errorf("model in not correct")
		}
	}

	return nil
}
