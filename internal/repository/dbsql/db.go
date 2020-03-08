package dbsql

import (
	"github.com/dickynovanto1103/User-Management-System/internal/model"
)

type DB interface {
	GetUser(username string) (model.User, error)
	UpdateNickname(nickname, username string) error
	UpdateProfile(profile, username string) error
	GetPassword(username string) (string, error)
}
