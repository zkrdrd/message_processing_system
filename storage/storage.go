package storage

import "messageProcessingSystem/model"

type Storage interface {
	SavePayment(*model.MessagePayment) error
	GetPaymentById(string) (*model.Payment, error)
}
