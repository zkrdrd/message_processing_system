package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"messageProcessingSystem/internal/model"
	"os"
)

// структура с путем сохраения файла базы данных
type DBLiteFiles struct {
	dbFile string //`default:"storage/lite/message.db"`
}

// заполнение структуры с путем сохраения файла базы данных
func NewDBLite(filePathToStorage string) *DBLiteFiles {
	return &DBLiteFiles{
		dbFile: filePathToStorage,
	}
}

// инициализация базы дынных
func (db *DBLiteFiles) InitLiteDatabase() {

	if _, err := os.Stat(db.dbFile); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(db.dbFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dbFileData.Close()

	if _, err = dbFileData.Exec(`CREATE TABLE IF NOT EXISTS payment 
		(id INTEGER PRIMARY KEY AUTOINCREMENT, 
		type_message TEXT NOT NULL, 
		uid_message TEXT NOT NULL UNIQUE, 
		address_from TEXT NULL, 
		address_to TEXT NULL, 
		payment INTEGER NULL);`); err != nil {
		log.Fatal(err)
	}
}

// сохранение и изменение данных в файл базы данных
func (db *DBLiteFiles) SavePayment(msg *model.Message) error {

	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	if _, err = dbFileData.ExecContext(context.Background(), `INSERT INTO payment (type_message, uid_message, address_from, address_to, payment) VALUES (?, ?, ?, ?, ?)
	ON CONFLICT (uid_message) DO UPDATE SET type_message = ?;`,
		msg.TypeMessage, msg.UidMessage, msg.AddressFrom, msg.AddressTo, msg.Payment, msg.TypeMessage); err != nil {
		return err
	}

	return nil
}
