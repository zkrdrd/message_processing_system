package sqlite

import (
	"database/sql"
	"errors"
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

	db.checkPaymentIsExist(dbFileData, msg)

	if _, err = dbFileData.Exec(`INSERT INTO payment (type_message, uid_message, address_from, address_to, payment, created_at) VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT DO UPDATE SET type_message = ?, modify_at = ? WHERE type_message='created';`,
		msg.TypeMessage, msg.UidMessage, msg.AddressFrom, msg.AddressTo, msg.Payment, time.Now().Format("01-02-2006 15:04:05"), msg.TypeMessage, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		return err
	}

	return nil
}

func (db *DBLite) GetPaymentById(uid string) (*model.GetedPayment, error) {
	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return nil, err
	}
	defer dbFileData.Close()

	gm := &model.GetedPayment{}

	err = dbFileData.QueryRow(`SELECT type_message, uid_message, address_from, address_to, payment, created_at, modify_at FROM payment WHERE uid_message = ?`, uid).Scan(&gm.TypeMessage, &gm.UidMessage, &gm.AddressFrom, &gm.AddressTo, &gm.Payment) //, &gm.CreatedAt, &gm.ModifyAt)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}
	return gm, nil
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
