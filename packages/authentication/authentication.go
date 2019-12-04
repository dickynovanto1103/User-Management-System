package authentication

import (
	"crypto/sha512"
	"dbutil"
	"encoding/hex"
	"errors"

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
	pass, err := dbutil.GetPassword(*username)
	if err != nil {
		return errors.New(dbutil.ErrorGetPassword)
	}

	if VerifyPassword(pass, *password) {
		return nil
	} else {
		return errors.New(ErrorNotAuthenticated)
	}
}
