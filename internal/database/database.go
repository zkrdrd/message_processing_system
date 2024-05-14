package database

import (
	"database/sql"
	"log"
)

func DBcreate() {
	db, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `payment` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `type_message` TEXT NOT NULL, `uid_message` TEXT NOT NULL UNIQUE, `address_from` TEXT NULL, `address_to` TEXT NULL, `payment` INTEGER NULL)")
	if err != nil {
		log.Fatal(err)
	}
}

func DBinsert() {
	db, err := sql.Open("sqlite3", "storage/lite/message.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into payment (type_message, uid_message, address_from, address_to, payment) values ($1, $2, $3, $4, $5)",
		"1", "2", "3", "4", 5)
	if err != nil {
		panic(err)
	}
}
