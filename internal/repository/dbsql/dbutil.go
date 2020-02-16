package dbsql

import (
	"database/sql"
	"log"

	"github.com/dickynovanto1103/User-Management-System/internal/service/config"
)

type DBImpl struct {
	db *sql.DB
}

func PrepareDB(config config.ConfigDB) (*sql.DB, error) {
	db, err := sql.Open(config.DriverName, config.Username+":"+config.Password+"@/"+config.DBName)
	if err != nil {
		log.Println("error opening DB: ", err)
		return nil, err
	}

	db.SetMaxOpenConns(MaxConnections)
	db.SetMaxIdleConns(MaxIdleConnections)
	return db, nil
}

func NewDBImpl(db *sql.DB) *DBImpl {
	return &DBImpl{
		db: db,
	}
}

func (impl *DBImpl) CloseDB() {
	impl.db.Close()
}

func (impl *DBImpl) PrepareQueryUser() (*sql.Stmt, error) {
	statementQueryUser, err := impl.db.Prepare("SELECT username, password, nickname, profileURL from " + tableName + " where username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	return statementQueryUser, nil
}

func (impl *DBImpl) PrepareUpdateNickname() (*sql.Stmt, error) {
	statementUpdateNickname, err := impl.db.Prepare("UPDATE " + tableName + " SET nickname = ? WHERE username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	return statementUpdateNickname, nil
}

func (impl *DBImpl) PrepareUpdateProfile() (*sql.Stmt, error) {
	statementUpdateProfile, err := impl.db.Prepare("UPDATE " + tableName + " SET profileURL = ? WHERE username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	return statementUpdateProfile, nil
}

func (impl *DBImpl) PrepareQueryPassword() (*sql.Stmt, error) {
	statementQueryPassword, err := impl.db.Prepare("SELECT password from " + tableName + " where username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	return statementQueryPassword, nil
}
