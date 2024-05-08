package messageProcessingSystem

type Message struct {
	TypeMessage string `json:"Type"`
	UidMessage  string `json:"Uid"`
	AddressFrom string `json:"AddressFrom,omitempty"`
	AddressTo   string `json:"AddressTo,omitempty"`
	Payment     int    `json:"Payment,omitempty"`
}

func (mes *Message) Processing(FileName string) {

}
