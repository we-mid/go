package db_sqlite

import (
	"database/sql"

	"github.com/we-task/Todo-as-a-Service/x/db"

	// _ "github.com/go-sql-driver/mysql"
	// sqlite3 driver for go using database/sql
	// _ "github.com/mattn/go-sqlite3"
	// pure-Go SQLite driver for Go (SQLite embedded)
	_ "github.com/glebarez/go-sqlite"
)

const (
	// driver = "mysql"
	// dsn    = "username:password@tcp(localhost:3306)/dbname"
	driver = "sqlite"
	// dsn    = "db.sqlite"
)

func NewDB(dsn, migrationDir string) (*sql.DB, error) {
	return db.NewDB(driver, dsn, migrationDir)
}
