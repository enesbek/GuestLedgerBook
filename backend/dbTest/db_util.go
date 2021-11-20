package dbTest

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"testing"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func mockDb(t testing.T) (*sql.DB, sqlmock.Sqlmock, *sqlx.DB){
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}
	sqlxDb := sqlx.NewDb(mockDb, "sqlmock")
	return mockDb, mock, sqlxDb
}
