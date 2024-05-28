package messageProcessingSystem

import (

	//"messageProcessingSystem/storage/memory"

	"errors"
	memory "messageProcessingSystem/storage"
	dblite "messageProcessingSystem/storage/lite"
	dbmemory "messageProcessingSystem/storage/memory"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zkrdrd/ConfigParser"
)

var ErrFieldIsEmpty = errors.New(`field 'UidMessage' is empty`)

// запись данных в структуру
func Reader(FileName string, Config any) error {
	if err := ConfigParser.Read(FileName, Config); err != nil {
		return err
	}
	return nil
}

// обработка json файлов
func Processing(msg *memory.Message) error {

	if msg.UidMessage == "" {
		return ErrFieldIsEmpty
	}

	lite := dblite.NewDBLite("storage/lite/message.db")
	lite.InitLiteDatabase()
	if err := lite.SavePayment(msg); err != nil {
		return err
	}

	if err := dbmemory.SavePayment(msg); err != nil {
		return err
	}

	return nil
}
