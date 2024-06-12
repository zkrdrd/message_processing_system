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

type GetMessage struct {
	Id           string
	Type_message string
	Address_from string
	Address_to   string
	Payment      int
	Created_at   string
	Modify_at    string
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

	db.checkPaymentIsExist(dbFileData, msg)

	if _, err = dbFileData.Exec(`INSERT INTO payment (type_message, uid_message, address_from, address_to, payment, created_at) VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT DO UPDATE SET type_message = ?, modify_at = ? WHERE type_message='created';`,
		msg.TypeMessage, msg.UidMessage, msg.AddressFrom, msg.AddressTo, msg.Payment, time.Now().Format("01-02-2006 15:04:05"), msg.TypeMessage, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		return err
	}

	return nil
}

func (db *DBLite) GetPaymentById(uid string) error {
	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	gm := &GetMessage{}

	err = dbFileData.QueryRow(`SELECT type_message, id, address_from, address_to, payment, created_at, modify_at FROM payment WHERE uid_message = ?`, uid).Scan(&gm.Type_message, &gm.Id, &gm.Address_from, &gm.Address_to, &gm.Payment, &gm.Created_at, &gm.Modify_at)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return err
	}
	fmt.Println(gm.Id, gm.Address_from, gm.Address_to, gm.Payment, gm.Created_at, gm.Modify_at)
	return nil
}

func (db *DBLite) checkPaymentIsExist(dbFileData *sql.DB, msg *model.Message) error {

	var (
		id           string
		type_message string
	)

	_ = dbFileData.QueryRow(`SELECT type_message, uid_message FROM payment WHERE type_message=? AND uid_message=?`, msg.TypeMessage, msg.UidMessage).Scan(&id, &type_message)

	if type_message == msg.TypeMessage && id == msg.UidMessage {
		log.Print("model is exist")
	}
	return nil
}
