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

func (p Payment) GetCreateAt() time.Time {
	time, _ := time.Parse(time.RFC3339, p.CreatedAt)
	return time
}

func (p Payment) GetUdatedAt() time.Time {
	time, _ := time.Parse(time.RFC3339, p.UpdatedAt)
	return time
}
