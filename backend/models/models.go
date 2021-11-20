package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	GuestModelImpl interface {
		FindByID(id string) Guest
		FindAll(offset int) []Guest
		CreateGuest(title string, message string) bool
	}

	Guest struct {
		ID   int    `json:"id" db:"id"`
		Title string `json:"title" db:"title"`
		Message string `json:"message" db:"message"`
	}

	GuestModel struct {
		db *sqlx.DB
	}
)

func NewGuestModel(db *sqlx.DB) *GuestModel {
	return &GuestModel{
		db: db,
	}
}

func (g *GuestModel) FindByID(id string) Guest {
	user := Guest{}
	g.db.Get(&user, "SELECT * FROM guestledger WHERE id = $1 LIMIT 1", id)
	return user
}

func (g *GuestModel) FindAll(offset int) []Guest {
	guestCount := 5
	rows, err := g.db.Query("SELECT * FROM guestledger ORDER BY id asc LIMIT $1 OFFSET $2", guestCount, offset*guestCount)
	if err != nil {
		panic(err)
	}

	guests := []Guest{}
	for rows.Next() {
		var g Guest
		if err := rows.Scan(&g.ID, &g.Title, &g.Message); err != nil {
			panic(err)
		}
		guests = append(guests, g)
	}

	return guests
}

func (g *GuestModel) CreateGuest(title string, message string) bool {
	sqlInsert := "INSERT INTO guestledger (title, message) VALUES($1, $2)"

	_, err := g.db.Exec(sqlInsert, title, message)
	if err != nil {
		return false
	}

	return true
}
