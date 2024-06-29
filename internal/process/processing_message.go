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

func NewMessagesProcessor(storage storage.Storage) *MessagesProcessor {
	return &MessagesProcessor{
		storage: storage,
	}
}

// обработка json файлов
func (mp *MessagesProcessor) PaymentProcessor(msg []byte) error {

	msgPayment := &model.MessagePayment{}

	if err := json.Unmarshal(msg, msgPayment); err != nil {
		return err
	}

	if err := msgPayment.Validate(); err != nil {
		return err
	}

	// 1. не обновлять платежи не имеющие статус created
	// 2. при создании платеж долже иметь amount > 0 должен иметь Address From и TO, не может придти id и статус больше 1 раза
	// 3. сделать функцию GetPaymentById - получение всего payment по id

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
