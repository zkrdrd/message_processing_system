package model

import (
	"errors"
	"fmt"
)

type Message struct {
	TypeMessage string `json:"TypeMessage"`
	UidMessage  string `json:"UidMessage"`
	AddressFrom string `json:"AddressFrom,omitempty"`
	AddressTo   string `json:"AddressTo,omitempty"`
	Payment     int    `json:"Payment,omitempty"`
}

var ErrFieldIsEmpty = errors.New(`field is empty`)

func (msg Message) Validate() error {
	if msg.UidMessage == "" {
		return fmt.Errorf("uid: %w", ErrFieldIsEmpty)
	}
	if msg.TypeMessage == "" {
		return fmt.Errorf("type message: %w", ErrFieldIsEmpty)
	}
	return nil
}

func (msg Message) ValidatePaymenIfNotExistInDB() error {
	if msg.AddressFrom == "" || msg.AddressTo == "" || msg.Payment <= 0 {
		return fmt.Errorf("model is not correct")
	}
	return nil
}

type GetedPayment struct {
	TypeMessage string
	UidMessage  string
	AddressFrom string
	AddressTo   string
	Payment     int
	//CreatedAt   string
	//ModifyAt    string
}
