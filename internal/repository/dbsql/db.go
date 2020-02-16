package dbsql

import (
	"database/sql"
)

type DB interface {
	PrepareQueryUser() (*sql.Stmt, error)
	PrepareUpdateNickname() (*sql.Stmt, error)
	PrepareUpdateProfile() (*sql.Stmt, error)
	PrepareQueryPassword() (*sql.Stmt, error)
	CloseDB()
}
