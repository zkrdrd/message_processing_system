package sqlite

import (
	"database/sql"
	"errors"
	"messageProcessingSystem/model"
	"os"
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
		amount INTEGER NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NULL);`); err != nil {
		return err
	}
	return nil
}

// сохранение и изменение данных в файл базы данных
func (db *DBLite) SavePayment(msg *model.MessagePayment) error {

	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return err
	}
	defer dbFileData.Close()

	if _, err = dbFileData.Exec(`
	INSERT INTO payment (type_message, uid_message, address_from, address_to, amount, created_at, updated_at) 
	VALUES (
		?, -- type_message
		?, -- uid_message
		?, -- address_from
		?, -- address_to
		?, -- amount
		?, -- created_at
		?) -- updated_at
	ON CONFLICT DO UPDATE SET 
		type_message = EXCLUDED.type_message, 
		updated_at = EXCLUDED.updated_at;`,
		msg.TypeMessage,
		msg.UidMessage,
		msg.AddressFrom,
		msg.AddressTo,
		msg.Amount,
		msg.CreatedAt,
		msg.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (db *DBLite) GetPaymentById(uid string) (*model.Payment, error) {
	dbFileData, err := sql.Open("sqlite3", db.dbFile)
	if err != nil {
		return nil, err
	}
	defer dbFileData.Close()

	gm := &model.Payment{}

	err = dbFileData.QueryRow(`
	SELECT type_message, uid_message, address_from, address_to, amount, created_at, updated_at 
	FROM payment WHERE uid_message = ?`, uid).Scan(
		&gm.TypeMessage,
		&gm.UidMessage,
		&gm.AddressFrom,
		&gm.AddressTo,
		&gm.Amount,
		&gm.CreatedAt,
		&gm.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrNotRows
		}
		return nil, err
	}
	return gm, nil
}
