package messageProcessingSystem

import (
	"github.com/zkrdrd/ConfigParser"
)

// обработка json файлов
func Processing(FileName string, config any) error {

	if err := ConfigParser.Read(FileName, config); err != nil {
		return err
	}
	return nil
}
