package db

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var createGuestTable string = `
DROP TABLE IF EXISTS guestledger;
CREATE TABLE guestledger
(
	id SERIAL PRIMARY KEY, 
	title TEXT,
	message TEXT
);
INSERT INTO guestledger(title, message) 
	VALUES
	    ('enes', 'hello'),
		('ahmet', 'hello 2');
`

func DBConnect() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "postgres://postgres:1234@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(createGuestTable)
	return db
}
