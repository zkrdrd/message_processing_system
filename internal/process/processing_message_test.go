package process_test

import (
	"fmt"
	"messageProcessingSystem/internal/model"
	"messageProcessingSystem/internal/process"
	"messageProcessingSystem/storage/memory"
	"testing"
)

func TestPaymentProcessor(t *testing.T) {
	t.Run(`test find min number`, func(t *testing.T) {
		testTable := TestParam

		pr := process.NewMessagesProcessor(memory.NewDatabase())
		for _, expect := range testTable {

			expect := expect

			if err := pr.PaymentProcessor([]byte(expect.Msg)); err != nil {
				if expect.Error != true {
					t.Error(fmt.Errorf(`result %v != %w`, expect.Error, err))
				}
			}

			/*if expect.Result != result {
				t.Error(fmt.Errorf(`result %s != %s`, expect.Result, result))
			}*/
		}
	})
}

var TestParam = []struct {
	Msg    string //*model.Message
	Result *model.GetedPayment
	Error  bool
}{
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "1A",
			"AddressFrom": "43245",
			"AddressTo":   "4124",
			"Payment":     5000
		}`,
		Result: &model.GetedPayment{},
		Error:  false,
	},
	{
		Msg: `{
			"TypeMessage": "processed",
			"UidMessage":  "1A"
		}`,
		Result: &model.GetedPayment{
			TypeMessage: "processed",
			UidMessage:  "1A",
			AddressFrom: "43234",
			AddressTo:   "4124",
			Payment:     5000,
		},
		Error: false,
	},
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "2A",
			"AddressFrom": "43224245",
			"AddressTo":   "41123424",
			"Payment":     500000
		}`,
		Result: &model.GetedPayment{
			TypeMessage: "created",
			UidMessage:  "2A",
			AddressFrom: "43224245",
			AddressTo:   "41123424",
			Payment:     500000,
		},
		Error: false,
	},
	{
		Msg: `{
			"TypeMessage": "created",
			"UidMessage":  "2A",
			"AddressFrom": "43224245",
			"AddressTo":   "41123424",
			"Payment":     500000
		}`,
		Result: &model.GetedPayment{
			TypeMessage: "created",
			UidMessage:  "2A",
			AddressFrom: "43224245",
			AddressTo:   "41123424",
			Payment:     500000,
		},
		Error: true,
	},
}
