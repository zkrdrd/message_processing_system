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
