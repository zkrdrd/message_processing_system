package messageProcessingSystem

import (

	//"messageProcessingSystem/storage/memory"

	database "messageProcessingSystem/storage/lite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zkrdrd/ConfigParser"
)

func Reader(FileName string, Config any) error {

	//var msg map[string]interface{}
	//var msg = &memory.Message{}
	if err := ConfigParser.Read(FileName, Config); err != nil {
		return err
	}
	Processing()
	return nil
}

// обработка json файлов
func Processing() error {

	//var data = map[string]memory.Message{}
	database.DBinsert() //memory.Message)
	//database.Testing()
	return nil
}
