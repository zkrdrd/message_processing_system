package memory

import (
	"messageProcessingSystem/model"
)

type DBMemory struct {
	inMemory map[string]model.Payment
}

func NewDatabase() *DBMemory {
	return &DBMemory{
		inMemory: make(map[string]model.Payment),
	}
}

// TODO:
// 1. статус платежа Бд не обновляется //
// 2. не проверяется если платеж уже прошел //

// сохранение данных в базу даннях в памяти
func (db *DBMemory) SavePayment(msg *model.Payment) error {

	db.inMemory[msg.UidMessage] = *msg
	return nil
}

func (db *DBMemory) GetPaymentById(id string) (*model.Payment, error) {
	if val, ok := db.inMemory[id]; ok {
		return &model.Payment{
			TypeMessage: val.TypeMessage,
			UidMessage:  val.UidMessage,
			AddressFrom: val.AddressFrom,
			AddressTo:   val.AddressTo,
			Amount:      val.Amount,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
		}, nil
	}
	return nil, model.ErrNotRows
}
