package messageProcessingSystem

import (
	database "messageProcessingSystem/storage/lite"

	//"messageProcessingSystem/storage/memory"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zkrdrd/ConfigParser"
)

// обработка json файлов
func Processing(FileName string, config any) error {
	//var msg = &memory.Message{}

	if err := ConfigParser.Read(FileName, config); err != nil {
		return err
	}

	database.DBinsert()
	//database.Testing()
	return nil
}
