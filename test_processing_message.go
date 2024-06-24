package process_test

import (
	"fmt"
	"messageProcessingSystem/internal/model"
	"messageProcessingSystem/internal/process"
	"messageProcessingSystem/storage/memory"
	"messageProcessingSystem/storage/sqlite"
	"sync"
	"testing"
)

func TestPaymentProcessor(t *testing.T) {
	t.Parallel()
	t.Run(`test find min number`, func(t *testing.T) {
		testTable := []struct {
			Msg    *model.Message
			Result *sqlite.GetMessage
			Error  bool
		}{
			{
				Msg: &model.Message{
					TypeMessage: "created",
					UidMessage:  "1A",
					AddressFrom: "43245",
					AddressTo:   "4124",
					Payment:     5000,
				},
				Result: &sqlite.GetMessage{},
				Error:  false,
			},
			{
				Msg: &model.Message{
					TypeMessage: "processed",
					UidMessage:  "1A",
				},
				Result: &sqlite.GetMessage{
					TypeMessage: "processed",
					UidMessage:  "1A",
					AddressFrom: "43234",
					AddressTo:   "4124",
					Payment:     5000,
				},
				Error: false,
			},
			{
				Msg: &model.Message{
					TypeMessage: "created",
					UidMessage:  "2A",
					AddressFrom: "43224245",
					AddressTo:   "41123424",
					Payment:     500000,
				},
				Result: &sqlite.GetMessage{
					TypeMessage: "created",
					UidMessage:  "2A",
					AddressFrom: "43224245",
					AddressTo:   "41123424",
					Payment:     500000,
				},
				Error: false,
			},
			{
				Msg: &model.Message{
					TypeMessage: "created",
					UidMessage:  "2A",
					AddressFrom: "43224245",
					AddressTo:   "41123424",
					Payment:     500000,
				},
				Result: &sqlite.GetMessage{
					TypeMessage: "created",
					UidMessage:  "2A",
					AddressFrom: "43224245",
					AddressTo:   "41123424",
					Payment:     500000,
				},
				Error: true,
			},
		}

		wg := sync.WaitGroup{}

		process.NewMessagesProcessor{memory.DBMemory{}}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наименьшего числа
				result := process.PaymentProcessor(expect.Msg, expect.Result, expect.Error)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %s != %s`, expect.Result, result))
				}
			}()
		}

		wg.Wait()
	})
}
