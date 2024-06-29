package model

import (
	"errors"
	"fmt"
)

const (
	TypeMessagePaymentCreated   = `created`
	TypeMessagePaymentProcessed = `processed`
	TypeMessagePaymentCanceled  = `canceled`
)

var (
	ErrFieldIsEmpty = errors.New(`field is empty`)
	ErrNotRows      = errors.New(`row is not found`)
)

type MessagePayment struct {
	TypeMessage string `json:"TypeMessage"`
	UidMessage  string `json:"UidMessage"`
	AddressFrom string `json:"AddressFrom,omitempty"`
	AddressTo   string `json:"AddressTo,omitempty"`
	Amount      int    `json:"Payment,omitempty"`
}

// Валидация обязательных полей
func (msg MessagePayment) Validate() error {
	if msg.UidMessage == "" {
		return fmt.Errorf("uid: %w", ErrFieldIsEmpty)
	}
	if msg.TypeMessage == "" {
		return fmt.Errorf("type message: %w", ErrFieldIsEmpty)
	}
	return nil
}
