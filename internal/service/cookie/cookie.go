package cookie

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter, sessionID string, duration time.Duration) {
	expiration := time.Now().Add(duration)
	cookie := http.Cookie{Name: "sessionID", Value: sessionID, Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println("getting error in retrieving cookie", err, cookie)
		return nil, errors.New("Forbidden")
	}

	return cookie, nil
}
