package dbutil

import (
	"config"
	"database/sql"
	"log"
	"user"
)

var db *sql.DB
var statementQueryUser, statementUpdateNickname, statementUpdateProfile, statementQueryPassword *sql.Stmt

const (
	MaxConnections      = 100
	MaxIdleConnections  = 100
	tableName           = "User"
	ErrorGetPassword    = "ErrorGetPassword"
	ErrorGetUser        = "ErrorGetUser"
	ErrorUpdateNickname = "ErrorUpdateNickname"
	ErrorUpdateProfile  = "ErrorUpdateProfile"
)

func PrepareDB(config c) {
	var err error
	db, err = sql.Open(config.DriverName, config.Username+":"+config.Password+"@/"+config.DBName)
	if err != nil {
		log.Println("error opening DB: ", err)
		return
	}
	db.SetMaxOpenConns(MaxConnections)
	db.SetMaxIdleConns(MaxIdleConnections)
}

func CloseDB() {
	db.Close()
}

func prepareQueryUser() {
	var err error
	statementQueryUser, err = db.Prepare("SELECT username, password, nickname, profileURL from " + tableName + " where username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
	}
}

func prepareUpdateNickname() {
	var err error
	statementUpdateNickname, err = db.Prepare("UPDATE " + tableName + " SET nickname = ? WHERE username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
	}
}

func prepareUpdateProfile() {
	var err error
	statementUpdateProfile, err = db.Prepare("UPDATE " + tableName + " SET profileURL = ? WHERE username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
	}
}

func prepareQueryPassword() {
	var err error
	statementQueryPassword, err = db.Prepare("SELECT password from " + tableName + " where username = ?")
	if err != nil {
		log.Println("error preparing statement: ", err)
	}
}

func GetPassword(username string) (string, error) {
	var pass string
	err := statementQueryPassword.QueryRow(username).Scan(&pass)
	return pass, err
}

func GetUser(username string) (user.User, error) {
	var data user.User
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
