package sql

import (
	"database/sql"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
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

func GetPassword(username string) (string, error) {
	var pass string
	err := statementQueryPassword.QueryRow(username).Scan(&pass)
	return pass, err
}

func GetUser(username string) (model.User, error) {
	var data model.User
	err := statementQueryUser.QueryRow(username).Scan(&data.Username, &data.Password, &data.Nickname, &data.ProfileURL)
	return data, err
}

func UpdateNickname(nickname string, username string) error {
	_, err := statementUpdateNickname.Exec(nickname, username)
	return err
}

func UpdateProfile(profile string, username string) error {
	_, err := statementUpdateProfile.Exec(profile, username)
	return err
}

func PrepareStatements() {
	prepareQueryUser()
	prepareUpdateNickname()
	prepareUpdateProfile()
	prepareQueryPassword()
}
