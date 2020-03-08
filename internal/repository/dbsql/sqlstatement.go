package dbsql

import (
	"database/sql"

	"github.com/dickynovanto1103/User-Management-System/internal/model"
)

type SQLStatements struct {
	StatementUpdateNickname *sql.Stmt
	StatementUpdateProfile  *sql.Stmt
	StatementQueryPassword  *sql.Stmt
	StatementQueryUser      *sql.Stmt
}

func NewSQLStatements(queryUser, queryPassword, updateNickname, updateProfile *sql.Stmt) *SQLStatements {
	return &SQLStatements{
		StatementUpdateNickname: updateNickname,
		StatementUpdateProfile:  updateProfile,
		StatementQueryPassword:  queryPassword,
		StatementQueryUser:      queryUser,
	}
}

func (s *SQLStatements) GetUser(username string) (model.User, error) {
	var data model.User
	err := s.StatementQueryUser.QueryRow(username).Scan(&data.Username, &data.Password, &data.Nickname, &data.ProfileURL)
	return data, err
}

func (s *SQLStatements) UpdateNickname(nickname, username string) error {
	_, err := s.StatementUpdateNickname.Exec(nickname, username)
	return err
}

func (s *SQLStatements) UpdateProfile(profile, username string) error {
	_, err := s.StatementUpdateProfile.Exec(profile, username)
	return err
}

func (s *SQLStatements) GetPassword(username string) (string, error) {
	var pass string
	err := s.StatementQueryPassword.QueryRow(username).Scan(&pass)
	return pass, err
}
