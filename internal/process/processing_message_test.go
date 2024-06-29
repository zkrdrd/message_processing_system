package process_test

import (
	"encoding/json"
	"fmt"
	"messageProcessingSystem/internal/process"
	"messageProcessingSystem/model"
	"messageProcessingSystem/storage/memory"
	"testing"
)

var storage = memory.NewDatabase()

func TestPaymentProcessor(t *testing.T) {
	t.Run(`test find min number`, func(t *testing.T) {
		testTable := TestParam

		pr := process.NewMessagesProcessor(storage)
		for _, expect := range testTable {

			msgPayment := &model.MessagePayment{}
			_ = json.Unmarshal([]byte(expect.Msg), msgPayment)

			expect := expect

			if err := pr.PaymentProcessor([]byte(expect.Msg)); err != nil {
				if expect.Error != true {
					t.Error(fmt.Errorf(`result %v != %w`, expect.Error, err))
				}
			}

			if model, err := storage.GetPaymentById(msgPayment.UidMessage); err != nil {
				if expect.Result != model {
					t.Error(fmt.Errorf(`result %v != %v`, expect.Result, model))
				}
			}
		}
	})
}

var TestParam = []struct {
	Msg    string //*model.Message
	Result *model.Payment
	Error  bool
}{
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "1A",
			"AddressFrom": "43245",
			"AddressTo":   "4124",
			"Amount":     5000
		}`,
		Result: &model.Payment{},
		Error:  false,
	},
	{
		Msg: `{
			"TypeMessage": "processed",
			"UidMessage":  "1A"
		}`,
		Result: &model.Payment{
			TypeMessage: "processed",
			UidMessage:  "1A",
			AddressFrom: "43234",
			AddressTo:   "4124",
			Amount:      5000,
		},
		Error: false,
	},
	{
		Msg: `{
			"TypeMessage": "canceled",
			"UidMessage":  "1A"
		}`,
		Result: &model.Payment{
			TypeMessage: "processed",
			UidMessage:  "1A",
			AddressFrom: "43234",
			AddressTo:   "4124",
			Amount:      5000,
		},
		Error: true,
	},
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "2A",
			"AddressFrom": "43224245",
			"AddressTo":   "41123424",
			"Amount":     500000
		}`,
		Result: &model.Payment{
			TypeMessage: "created",
			UidMessage:  "2A",
			AddressFrom: "43224245",
			AddressTo:   "41123424",
			Amount:      500000,
		},
		Error: false,
	},
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "2A",
			"AddressFrom": "43224245",
			"AddressTo":   "41123424",
			"Amount":     500000
		}`,
		Result: &model.Payment{
			TypeMessage: "created",
			UidMessage:  "2A",
			AddressFrom: "43224245",
			AddressTo:   "41123424",
			Amount:      500000,
		},
		Error: true,
	},
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "3A",
			"AddressFrom": "1234",
			"AddressTo":   "1234",
			"Amount":     100
		}`,
		Result: &model.Payment{
			TypeMessage: "created",
			UidMessage:  "3A",
			AddressFrom: "1234",
			AddressTo:   "1234",
			Amount:      100,
		},
		Error: false,
	},
	{
		Msg: `{
			"TypeMessage": "canceled",
			"UidMessage":  "3A",
			"AddressFrom": "1234",
			"AddressTo":   "1234",
			"Amount":     100
		}`,
		Result: &model.Payment{
			TypeMessage: "canceled",
			UidMessage:  "3A",
			AddressFrom: "1234",
			AddressTo:   "1234",
			Amount:      100,
		},
		Error: false,
	},
	{
		Msg: `{
			"TypeMessage": "processed",
			"UidMessage":  "3A",
			"AddressFrom": "1234",
			"AddressTo":   "1234",
			"Amount":     100
		}`,
		Result: &model.Payment{
			TypeMessage: "processed",
			UidMessage:  "3A",
			AddressFrom: "1234",
			AddressTo:   "1234",
			Amount:      100,
		},
		Error: true,
	},
	{
		Msg: `{
			"TypeMessage": "processed",
			"UidMessage":  "4A",
			"AddressFrom": "1234",
			"AddressTo":   "1234",
			"Amount":     1000
		}`,
		Result: &model.Payment{
			TypeMessage: "processed",
			UidMessage:  "3A",
			AddressFrom: "1234",
			AddressTo:   "1234",
			Amount:      1000,
		},
		Error: true,
	},
}
