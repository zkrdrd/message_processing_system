package process

import (
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
func (mp *MessagesProcessor) PaymentProcessor(msg []byte) error {

	msgPayment := &model.MessagePayment{}

	if err := json.Unmarshal(msg, msgPayment); err != nil {
		return err
	}

	// Валидация обязытельных полей
	if err := msgPayment.Validate(); err != nil {
		return err
	}

	// Получение данных из базы по id
	// Если данных нет то проверяем поля которые должны быть если сообщение новое
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

			// Сохрание данных
			if err := mp.storage.SavePayment(payments); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// Пhоверяем статус платежа если он created то обновляем данные
	if err = CompareOldAndNewStatePayment(msgPayment, payment); err != nil {
		return err
	}

	payment.TypeMessage = msgPayment.TypeMessage
	payment.UpdatedAt = model.SetDateTime()

	// Сохрание данных
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
