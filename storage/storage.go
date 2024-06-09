package storage

import "messageProcessingSystem/internal/model"

type Storage interface {
	SavePayment(*model.Message) error
	GetPaymentById(string) error
	CheckDatabaseAndModelIsCorrect(*model.Message) error
}
