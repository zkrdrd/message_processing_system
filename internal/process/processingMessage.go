package messageProcessingSystem

import (

	//"messageProcessingSystem/storage/memory"

	memory "messageProcessingSystem/storage"
	database "messageProcessingSystem/storage/lite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zkrdrd/ConfigParser"
)

// запись данных в структуру
func Reader(FileName string, Config any) error {
	if err := ConfigParser.Read(FileName, Config); err != nil {
		return err
	}
	return nil
}

// обработка json файлов
func Processing(msg *memory.Message) error {
	database.SavePayment(msg)
	return nil
}
