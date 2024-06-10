package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"messageProcessingSystem/internal/model"
	"os"
	"time"
)

// структура с путем сохраения файла базы данных
type DBLite struct {
	dbFile string //`default:"storage/lite/message.db"`
}

// заполнение структуры с путем сохраения файла базы данных
func NewDatabase(filePathToStorage string) *DBLite {
	return &DBLite{
		dbFile: filePathToStorage,
	}
}

// инициализация базы дынных
func (db *DBLite) InitLiteDatabase() error {

	if _, err := os.Stat(db.dbFile); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(db.dbFile)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	if _, err = dbFileData.Exec(`CREATE TABLE IF NOT EXISTS payment 
		(id INTEGER PRIMARY KEY AUTOINCREMENT, 
		type_message TEXT NOT NULL, 
		uid_message TEXT NOT NULL UNIQUE, 
		address_from TEXT NULL, 
		address_to TEXT NULL, 
		payment INTEGER NULL,
		created_at TEXT NOT NULL,
		modify_at TEXT NULL);`); err != nil {
		return err
	}
	return nil
}

// сохранение и изменение данных в файл базы данных
func (db *DBLite) SavePayment(msg *model.Message) error {

	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	if _, err = dbFileData.Exec(`INSERT INTO payment (type_message, uid_message, address_from, address_to, payment, created_at) VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT DO UPDATE SET type_message = ?, modify_at = ? WHERE type_message='created';`,
		msg.TypeMessage, msg.UidMessage, msg.AddressFrom, msg.AddressTo, msg.Payment, time.Now().Format("01-02-2006 15:04:05"), msg.TypeMessage, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		return err
	}

	return nil
}

func (db *DBLite) GetPaymentById(id string) error {
	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	row, err := dbFileData.Query(`SELECT type_message, uid_message, address_from, address_to, payment, created_at, modify_at FROM payment WHERE uid_message=?`, id)
	if err != nil {
		return err
	}
	for row.Next() {
		var (
			id           string
			type_message string
			address_from string
			address_to   string
			payment      int
			created_at   string
			modify_at    string
		)
		row.Scan(&type_message, &id, &address_from, &address_to, &payment, &created_at, &modify_at)
		fmt.Println(id, type_message, address_from, address_to, payment, created_at, modify_at)
	}
	return nil
}

// проверка элементов базы данных
func (db *DBLite) CheckDatabaseAndModelIsCorrect(msg *model.Message) error {
	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	row, err := dbFileData.Query(`SELECT type_message, uid_message FROM payment WHERE type_message=? AND uid_message=?`, msg.TypeMessage, msg.UidMessage)
	if err != nil {
		return err
	}
	for row.Next() {
		var (
			id           string
			type_message string
		)
		row.Scan(&type_message, &id)

		if type_message == msg.TypeMessage && id == msg.UidMessage {
			return fmt.Errorf("model is exist")
		}
	}

	var id string
	err = dbFileData.QueryRow(`SELECT uid_message FROM payment WHERE uid_message = ?`, msg.UidMessage).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return err
	}

	return nil
}
