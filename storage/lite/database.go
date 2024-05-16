package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

//var msg = &memory.Message{}

func NewSqlite() {
	if _, err := os.Stat("storage/lite/message.db"); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create("storage/lite/message.db")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		DBcreate()
	}

}

func DBcreate() {
	db, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `payment` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `type_message` TEXT NOT NULL, `uid_message` TEXT NOT NULL UNIQUE, `address_from` TEXT NULL, `address_to` TEXT NULL, `payment` INTEGER NULL);")
	if err != nil {
		log.Fatal(err)
	}
}

func DBinsert() error {
	db, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		return err
	}
	defer db.Close()

	/*_, err = db.Exec("INSERT INTO payment (type_message, uid_message, address_from, address_to, payment) VALUES (?, ?, ?, ?, ?);",
	msg.TypeMessage, msg.UidMessage, msg.AddressFrom, msg.AddressTo, msg.Payment)*/

	/*prp, _ := db.Prepare("INSERT INTO payment (type_message, uid_message, address_from, address_to, payment) VALUES (?, ?, ?, ?, ?);")
	_, err = prp.Exec(2, 3, 4, 5, 6)*/

	_, err = db.ExecContext(context.Background(), "INSERT INTO payment (type_message, uid_message, address_from, address_to, payment) VALUES (?, ?, ?, ?, ?);",
		3, 4, 5, 6, 7)
	if err != nil {
		return err
	}
	return nil
}

func Testing() error {
	db, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, _ = db.ExecContext(context.Background(), "INSERT INTO payment (type_message, uid_message, address_from, address_to, payment) VALUES (?, ?, ?, ?, ?);",
		1, 2, 3, 4, 5)
	//row, _ := db.Query("SELECT * FROM payment;")
	exist, err := db.QueryContext(context.Background(), "SELECT type_message FROM payment WHERE uid_message = '?';", 3)
	if exist != nil {
		fmt.Println("asd")
	}
	if err != nil {
		return err
	}
	/*var (
		TypeMessage string
		UidMessage  string
		AddressFrom string
		AddressTo   string
		Payment     int
	)
	for row.Next() {
		row.Scan(&TypeMessage, &UidMessage, &AddressFrom, &AddressTo, &Payment)
		fmt.Println(TypeMessage, UidMessage, AddressFrom, AddressTo, Payment)
	}*/
	return nil
}
