package storage

import "messageProcessingSystem/model"

type Storage interface {
	SavePayment(*model.Payment) error
	GetPaymentById(string) (*model.Payment, error)
}
