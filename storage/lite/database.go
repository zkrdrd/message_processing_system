package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
	memory "messageProcessingSystem/storage"
	"os"
)

func InitDatabase() {
	if _, err := os.Stat("storage/lite/message.db"); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create("storage/lite/message.db")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		CreateDB()
	}
}

func CreateDB() {
	dbFile, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbFile.Close()

	_, err = dbFile.Exec(`CREATE TABLE IF NOT EXISTS payment 
		(id INTEGER PRIMARY KEY AUTOINCREMENT, 
		type_message TEXT NOT NULL, 
		uid_message TEXT NOT NULL UNIQUE, 
		address_from TEXT NULL, 
		address_to TEXT NULL, 
		payment INTEGER NULL);`)
	if err != nil {
		log.Fatal(err)
	}
}

func SavePayment(msg *memory.Message) error {
	dbFile, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		return err
	}
	defer dbFile.Close()

	if msg.UidMessage == "" {
		return err
	}

	_, err = dbFile.ExecContext(context.Background(), `INSERT INTO payment (type_message, uid_message, address_from, address_to, payment) VALUES (?, ?, ?, ?, ?)
	ON CONFLICT (uid_message) DO UPDATE SET type_message = ?;`,
		msg.TypeMessage, msg.UidMessage, msg.AddressFrom, msg.AddressTo, msg.Payment, msg.TypeMessage)
	if err != nil {
		return err
	}
	return nil
}
