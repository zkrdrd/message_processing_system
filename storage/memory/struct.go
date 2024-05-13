package memory

type Message struct {
	TypeMessage string `json:"TypeMesage"`
	UidMessage  string `json:"UidMessage"`
	AddressFrom string `json:"AddressFrom,omitempty"`
	AddressTo   string `json:"AddressTo,omitempty"`
	Payment     int    `json:"Payment,omitempty"`
}
