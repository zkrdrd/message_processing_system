package process

import (

	//"messageProcessingSystem/storage/memory"

	"encoding/json"
	"fmt"
	"messageProcessingSystem/model"
	"messageProcessingSystem/storage"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrAmountLessOne      = fmt.Errorf(`field amount is less 1`)
	ErrPaymentIsExist     = fmt.Errorf(`payment is exist`)
	ErrPaymentISCompleted = fmt.Errorf(`payment is completed`)
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

	strc, err := mp.storage.GetPaymentById(msgPayment.UidMessage)
	if err == nil {
		if err = ValidateStrc(msgPayment, strc); err != nil {
			return err
		}
	}

	if err := mp.storage.SavePayment(msgPayment); err != nil {
		return err
	}

	strc1, _ := mp.storage.GetPaymentById(msgPayment.UidMessage)
	fmt.Println(strc1.TypeMessage)

	return nil
}

func ValidateStrc(msgPayment *model.MessagePayment, strc *model.Payment) error {
	if strc.TypeMessage == model.TypeMessagePaymentCreated {
		if strc.TypeMessage == msgPayment.TypeMessage {
			return ErrPaymentIsExist
		}
	}
	if strc.TypeMessage == model.TypeMessagePaymentProcessed || strc.TypeMessage == model.TypeMessagePaymentCanceled {
		return ErrPaymentISCompleted
	}
	if strc.Amount == 0 && msgPayment.Amount < 1 {
		return ErrAmountLessOne
	}
	return nil
}
