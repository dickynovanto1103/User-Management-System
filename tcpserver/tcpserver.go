package main

import (
	"encoding/gob"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/request"
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/stringutil"

	"github.com/dickynovanto1103/User-Management-System/internal/config"

	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"

	"github.com/dickynovanto1103/User-Management-System/internal/authentication"

	"github.com/dickynovanto1103/User-Management-System/internal/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/profile"
)

func handleAuthentication(conn net.Conn, req request.Request) {
	username := req.Data[user.CodeUsername].(string)
	password := req.Data[user.CodePassword].(string)
	err := authentication.Authenticate(&username, &password)
	if err != nil {
		if err.Error() == authentication.ErrorNotAuthenticated {
			conn.Write([]byte(authentication.ErrorNotAuthenticated + "\n"))
		} else {
			conn.Write([]byte(err.Error() + "\n"))
		}
	} else {
		sessionID := stringutil.CreateRandomString(32)
		redisutil.Set(sessionID, username, 5*time.Hour)
		conn.Write([]byte(sessionID + "\n"))
	}
}

func getUserDataFromCache(username string) (user.User, error) {
	var data user.User
	var err error
	data.Username, err = redisutil.Get(username + user.CodeUsername)
	if err != nil {
		return user.User{}, err
	}
	data.Nickname, err = redisutil.Get(username + user.CodeNickname)
	if err != nil {
		return user.User{}, err
	}
	data.ProfileURL, err = redisutil.Get(username + user.CodeProfile)
	if err != nil {
		return user.User{}, err
	}
	return data, nil
}

func setUserDataCache(data user.User) {
	redisutil.Set(data.Username+user.CodeUsername, data.Username, 5*time.Minute)
	redisutil.Set(data.Username+user.CodeNickname, data.Nickname, 5*time.Minute)
	redisutil.Set(data.Username+user.CodeProfile, data.ProfileURL, 5*time.Minute)
}

func getUserData(username string) (user.User, error) {
	data, err := getUserDataFromCache(username)
	if err != nil {
		log.Println("fail to get data from cache", err)
		data, err = dbutil.GetUser(username)
		if err != nil {
			log.Println("fail to get user data from database", err)
			return user.User{}, err
		}
		log.Println("data from database: ", data)
		setUserDataCache(data)
	}
	return data, err
}

func responseForbidden(conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	var mapper = make(map[string]interface{})
	mapper[response.ResponseCode] = response.ResponseKeyForbidden
	resp := response.Response{ResponseID: response.ResponseForbidden, Data: mapper}
	err := encoder.Encode(resp)
	if err != nil {
		log.Println("error in encoding response: ", err)
	}
	log.Println("handle info response: ", resp)
}

func responseUser(conn net.Conn, data user.User) {
	encoder := gob.NewEncoder(conn)
	var mapper = make(map[string]interface{})
	mapper[user.CodeUser] = data
	resp := response.Response{ResponseID: response.ResponseOK, Data: mapper}
	err := encoder.Encode(resp)
	log.Println("handle info response: ", resp)
	if err != nil {
		log.Println("error in encoding response: ", err)
	}
}

func handleInfo(conn net.Conn, req request.Request) {
	cookie := req.Data[user.CodeCookie].(http.Cookie)
	username, err := redisutil.Get(cookie.Value)
	if err != nil {
		log.Println("error retrieving user data from cookie by redis: ", err)
		responseForbidden(conn)
		return
	}
	data, err := getUserData(username)
	if err != nil {
		log.Println("fail to get user data", err)
		responseForbidden(conn)
	} else {
		responseUser(conn, data)
	}

}

func sendResponseBack(conn net.Conn, username string, err error) {
	if err != nil {
		log.Println("error update statement", err)
		responseForbidden(conn)
	} else {
		data, err := getUserData(username)
		if err != nil {
			log.Println("error getting user data from database: ", err)
			return
		}
		responseUser(conn, data)
	}
}

func handleUpdateNickname(conn net.Conn, req request.Request) {
	nickname := req.Data[user.CodeNickname].(string)
	cookie := req.Data[user.CodeCookie].(http.Cookie)
	username, err := redisutil.Get(cookie.Value)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		responseForbidden(conn)
		return
	}

	err = dbutil.UpdateNickname(nickname, username)
	if err != nil {
		log.Println(dbutil.ErrorUpdateProfile + " " + err.Error())
		return
	}
	user, err := dbutil.GetUser(username)
	setUserDataCache(user)

	sendResponseBack(conn, username, err)
}

func handleUpdateProfile(conn net.Conn, req request.Request) {
	profile := req.Data[user.CodeProfile].(string)
	cookie := req.Data[user.CodeCookie].(http.Cookie)
	username, err := redisutil.Get(cookie.Value)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		responseForbidden(conn)
		return
	}

	err = dbutil.UpdateProfile(profile, username)
	if err != nil {
		log.Println(dbutil.ErrorUpdateProfile + " " + err.Error())
		return
	}
	user, err := dbutil.GetUser(username)
	setUserDataCache(user)

	sendResponseBack(conn, username, err)
}

func handleConnection(conn net.Conn) {
	for {
		decoder := gob.NewDecoder(conn)
		var req request.Request
		err := decoder.Decode(&req)
		if err != nil {
			log.Println("error in decoding in handling connection", err)
			break
		}
		if req.RequestID == request.RequestLogin {
			handleAuthentication(conn, req)
		} else if req.RequestID == request.RequestUserInfo {
			handleInfo(conn, req)
		} else if req.RequestID == request.RequestUpdateNickname {
			handleUpdateNickname(conn, req)
		} else if req.RequestID == request.RequestUpdateProfile {
			handleUpdateProfile(conn, req)
		}
	}
}

func main() {
	configDB := config.LoadConfigDB("config/configDB.json")
	configRedis := config.LoadConfigRedis("config/configRedis.json")

	redisutil.CreateRedisClient(configRedis)

	defer profile.Start().Stop()
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Println("error found in listening: ", err)
	}

	dbutil.PrepareDB(configDB)
	defer dbutil.CloseDB()
	dbutil.PrepareStatements()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("error found in accepting connection: ", err)
			continue
		}
		go handleConnection(connection)
	}
}
