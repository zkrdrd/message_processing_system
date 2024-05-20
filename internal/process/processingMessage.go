package messageProcessingSystem

import (

	//"messageProcessingSystem/storage/memory"
	database "messageProcessingSystem/storage/lite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zkrdrd/ConfigParser"
)

// обработка json файлов
func Processing(FileName string, config any) error {
	//var msg = &memory.Message{}
	//var data = map[string]memory.Message{}
	if err := ConfigParser.Read(FileName, config); err != nil {
		return err
	}
	database.DBinsert() //memory.Message)
	//database.Testing()
	return nil
}
