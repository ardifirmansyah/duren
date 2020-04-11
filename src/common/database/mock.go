package database

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type (
	MockDatabase struct {
		mockDBConn *sql.DB
		mocker     sqlmock.Sqlmock
	}
)
