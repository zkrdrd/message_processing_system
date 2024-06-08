package messageProcessingSystem

import (

	//"messageProcessingSystem/storage/memory"

	"errors"
	"messageProcessingSystem/internal/model"
	dblite "messageProcessingSystem/storage/lite"

	_ "github.com/mattn/go-sqlite3"
)

var ErrFieldIsEmpty = errors.New(`field 'UidMessage' is empty`)

type Storage struct {
	storageFIle string
}

func NewStorage(filepath string) *Storage {
	return &Storage{
		storageFIle: filepath,
	}
}

// обработка json файлов
func Processing(msg *model.Message) error {

	if msg.UidMessage == "" {
		return ErrFieldIsEmpty
	}

	lite := dblite.NewDBLite("storage/lite/message.db")
	lite.InitLiteDatabase()
	if err := lite.SavePayment(msg); err != nil {
		return err
	}

	/*mem := dbmemory.NewDatabase()
	if err := mem.SavePayment(msg); err != nil {
		return err
	}*/

	return nil
}
