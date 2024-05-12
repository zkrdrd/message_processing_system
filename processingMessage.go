package messageProcessingSystem

import (
	"encoding/json"
	"os"
)

func Processing(FileName string, config any) error {
	readData, err := os.ReadFile(FileName)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(readData, config); err != nil {
		return err
	}

	return nil
}
