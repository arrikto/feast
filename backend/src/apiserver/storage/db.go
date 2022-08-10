package storage

import (
	"database/sql"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	sqlite3 "github.com/mattn/go-sqlite3"
)

// DB a struct wrapping plain sql library with SQL dialect, to solve any feature
// difference between MySQL, which is used in production, and Sqlite, which is used
// for unit testing.
type DB struct {
	*sql.DB
	SQLDialect
}

// NewDB creates a DB
func NewDB(db *sql.DB, dialect SQLDialect) *DB {
	return &DB{db, dialect}
}

// SQLDialect abstracts common sql queries which vary in different dialect.
// It is used to bridge the difference between mysql (production) and sqlite
// (test).
type SQLDialect interface {
	// Check whether the error is a SQL duplicate entry error or not
	IsDuplicateError(err error) bool
}

// MySQLDialect implements SQLDialect with mysql dialect implementation.
type MySQLDialect struct{}

func (d MySQLDialect) IsDuplicateError(err error) bool {
	sqlError, ok := err.(*mysql.MySQLError)
	return ok && sqlError.Number == mysqlerr.ER_DUP_ENTRY
}

// SQLiteDialect implements SQLDialect with sqlite dialect implementation.
type SQLiteDialect struct{}

func (d SQLiteDialect) IsDuplicateError(err error) bool {
	sqlError, ok := err.(sqlite3.Error)
	return ok && sqlError.Code == sqlite3.ErrConstraint
}

func NewMySQLDialect() MySQLDialect {
	return MySQLDialect{}
}

func NewSQLiteDialect() SQLiteDialect {
	return SQLiteDialect{}
}
