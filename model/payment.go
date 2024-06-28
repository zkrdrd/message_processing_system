package model

import "time"

const FormatDateTime = `01-02-2006 15:04:05`

type Payment struct {
	TypeMessage string
	UidMessage  string
	AddressFrom string
	AddressTo   string
	Amount      int
	CreatedAt   string
	UpdatedAt   string
}

func GetCreateAt() string {
	return time.Now().Format(FormatDateTime)
}

func GetUdatedAt() string {
	return time.Now().Format(FormatDateTime)
}
