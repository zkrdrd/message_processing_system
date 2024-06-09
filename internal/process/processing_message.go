package process

import (

	//"messageProcessingSystem/storage/memory"

	"encoding/json"
	"messageProcessingSystem/internal/model"
	"messageProcessingSystem/storage"

	_ "github.com/mattn/go-sqlite3"
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

	payment := &model.Message{}
	if err := json.Unmarshal(msg, payment); err != nil {
		return err
	}

	if err := payment.Validate(); err != nil {
		return err
	}

	if err := mp.storage.SavePayment(payment); err != nil {
		return err
	}

	if err := mp.storage.GetPaymentById(payment.UidMessage); err != nil {
		return err
	}

	return nil
}
