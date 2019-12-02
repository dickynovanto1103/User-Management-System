package user

//User user contains username and password.
type User struct {
	UserID     int
	Username   string
	Password   string
	Nickname   string
	ProfileURL string
}

const (
	CodeUsername  = "username"
	CodePassword  = "password"
	CodeUser      = "user"
	CodeProfile   = "profile"
	CodeNickname  = "nickname"
	CodeCookie    = "cookie"
	CodeSessionID = "sessionID"
)
