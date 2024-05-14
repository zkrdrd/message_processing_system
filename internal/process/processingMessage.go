package messageProcessingSystem

import (
	"log"
	"messageProcessingSystem/internal/database"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zkrdrd/ConfigParser"
)

func init() {
	f, err := os.Create("storage/lite/message.db")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	database.DBcreate()
}

// обработка json файлов
func Processing(FileName string, config any) error {

	if err := ConfigParser.Read(FileName, config); err != nil {
		return err
	}
	database.DBinsert()

	return nil
}
