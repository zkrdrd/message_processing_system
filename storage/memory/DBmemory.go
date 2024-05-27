package memory

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	message "messageProcessingSystem/storage"
)

func CreateDB() {
	dbFile, err := sql.Open("sqlite3", "file:memorybase?mode=memory&cache=shared")
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

func SavePayment(msg *message.Message) error {
	dbFile, err := sql.Open("sqlite3", "file:memorybase?mode=memory&cache=shared")
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

	rows := 0
	dbFile.QueryRow("SELECT COUNT(*) FROM payment;").Scan(&rows)
	fmt.Println(rows)

	return nil
}
