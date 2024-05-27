package messageProcessingSystem

import (

	//"messageProcessingSystem/storage/memory"

	memory "messageProcessingSystem/storage"
	dblite "messageProcessingSystem/storage/lite"
	dbmemory "messageProcessingSystem/storage/memory"

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

	dblite.InitDatabase()
	dbmemory.CreateDB()

	if err := dblite.SavePayment(msg); err != nil {
		return err
	}
	if err := dbmemory.SavePayment(msg); err != nil {
		return err
	}
	return nil
}
