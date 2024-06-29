package process

import (

	//"messageProcessingSystem/storage/memory"

	"encoding/json"
	"errors"
	"messageProcessingSystem/model"
	"messageProcessingSystem/storage"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrAmountLessOne      = errors.New(`field amount is less 1`)
	ErrPaymentIsExist     = errors.New(`payment is exist`)
	ErrPaymentISCompleted = errors.New(`payment is completed`)
	ErrAddressFromIsEmpty = errors.New(`in model payment address from is empry`)
	ErrAddressToIsEmpty   = errors.New(`in model payment address from is empry`)
)

type MessagesProcessor struct {
	storage storage.Storage
}

// определение хранилица
func NewMessagesProcessor(storage storage.Storage) *MessagesProcessor {
	return &MessagesProcessor{
		storage: storage,
	}
}

// 1. Обработка сообщения
// 2. Валидация обязытельных полей
// 3. Получение данных из базы по id
// 3.1. Если данных нет то проверяем поля которые должны быть если сообщение новое
// 3.2. Сохраняем данные
// 4. Если данные в базе присутствуют
// 4.1. ПРоверяем статус платежа
// 4.2 Если он created то обновляем данные
func (mp *MessagesProcessor) PaymentProcessor(msg []byte) error {

	msgPayment := &model.MessagePayment{}

	if err := json.Unmarshal(msg, msgPayment); err != nil {
		return err
	}

	if err := msgPayment.Validate(); err != nil {
		return err
	}

	payment, err := mp.storage.GetPaymentById(msgPayment.UidMessage)
	if err != nil {
		if err == model.ErrNotRows {

			if msgPayment.AddressFrom == "" {
				return ErrAddressFromIsEmpty
			}
			if msgPayment.AddressTo == "" {
				return ErrAddressFromIsEmpty
			}
			if msgPayment.Amount < 1 {
				return ErrAmountLessOne
			}

			payments := &model.Payment{
				TypeMessage: msgPayment.TypeMessage,
				UidMessage:  msgPayment.UidMessage,
				AddressFrom: msgPayment.AddressFrom,
				AddressTo:   msgPayment.AddressTo,
				Amount:      msgPayment.Amount,
				CreatedAt:   model.SetDateTime(),
				UpdatedAt:   model.SetDateTime(),
			}

			if err := mp.storage.SavePayment(payments); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	if err = CompareOldAndNewStatePayment(msgPayment, payment); err != nil {
		return err
	}

	payment.TypeMessage = msgPayment.TypeMessage
	payment.UpdatedAt = model.SetDateTime()

	if err := mp.storage.SavePayment(payment); err != nil {
		return err
	}
	return nil
}

// Проверяем статус платежа
func CompareOldAndNewStatePayment(msgPayment *model.MessagePayment, payment *model.Payment) error {
	if payment.TypeMessage == model.TypeMessagePaymentCreated {
		if payment.TypeMessage == msgPayment.TypeMessage {
			return ErrPaymentIsExist
		}
	}
	if payment.TypeMessage == model.TypeMessagePaymentProcessed ||
		payment.TypeMessage == model.TypeMessagePaymentCanceled {
		return ErrPaymentISCompleted
	}
	return nil
}
