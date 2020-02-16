package authentication

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/dickynovanto1103/User-Management-System/container"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"

	"golang.org/x/crypto/pbkdf2"
)

const (
	ErrorNotAuthenticated = "Not Authenticated"
)

func VerifyPassword(storedPassword string, inputPassword string) bool {
	salt := storedPassword[:64]
	storedPassword = storedPassword[64:]
	hashedPassword := pbkdf2.Key([]byte(inputPassword), []byte(salt), 1, 64, sha512.New)
	strHashedPassword := hex.EncodeToString(hashedPassword)
	return storedPassword == strHashedPassword
}

func Authenticate(username *string, password *string) error {
	pass, err := getPassword(*username)
	if err != nil {
		return errors.New(dbsql.ErrorGetPassword)
	}

	if VerifyPassword(pass, *password) {
		return nil
	} else {
		return errors.New(ErrorNotAuthenticated)
	}
}

func getPassword(username string) (string, error) {
	var pass string
	err := container.StatementQueryPassword.QueryRow(username).Scan(&pass)
	return pass, err
}
