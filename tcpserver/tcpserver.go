package main

import (
	"context"
	"encoding/gob"
	"google.golang.org/grpc"
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

	pb "github.com/dickynovanto1103/User-Management-System/proto"
)


type server struct {
	pb.UnimplementedUserDataServiceServer
}

func (s *server) SendRequest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	requestID := in.GetRequestID()
	mapper := in.GetMapper()
	response := handleRequest(requestID, mapper)
	newResponse := &pb.Response{
		ResponseID:           int32(response.ResponseID),
		Mapper:               response.Data,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return newResponse, nil
}

func handleAuthentication(mapper map[string]string) response.Response {
	username := mapper[user.CodeUsername]
	password := mapper[user.CodePassword]
	err := authentication.Authenticate(&username, &password)
	mapperResp := make(map[string]string)
	if err != nil {
		if err.Error() == authentication.ErrorNotAuthenticated {
			mapperResp[response.ResponseCode] = authentication.ErrorNotAuthenticated
		} else {
			mapperResp[response.ResponseCode] = err.Error()
		}
	} else {
		sessionID := stringutil.CreateRandomString(32)
		redisutil.Set(sessionID, username, 5*time.Hour)
		mapperResp[response.ResponseCode] = sessionID
	}
	return response.Response{ResponseID: response.ResponseOK, Data: mapperResp}
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

func getMapFromUser(data user.User) map[string]string {
	var mapper = make(map[string]string)
	mapper[user.CodeUsername] = data.Username
	mapper[user.CodeNickname] = data.Nickname
	mapper[user.CodeProfile] = data.ProfileURL
	return mapper
}

func handleInfo(mapper map[string]string) response.Response {
	userIDFromCookie := mapper[user.CodeCookie]
	username, err := redisutil.Get(userIDFromCookie)
	if err != nil {
		log.Println("error retrieving user data from cookie by redis: ", err)
		return responseForbidden()
	}
	data, err := getUserData(username)
	if err != nil {
		log.Println("fail to get user data", err)
		return responseForbidden()
	} else {
		return responseUser(data)
	}

}

func sendResponseBack(username string, err error) response.Response {
	if err != nil {
		log.Println("error update statement", err)
		return responseForbidden()
	} else {
		data, err := getUserData(username)
		if err != nil {
			log.Println("error getting user data from database: ", err)
			return responseForbidden()
		}
		return responseUser(data)
	}
}

func handleUpdateNickname(mapper map[string]string) response.Response {
	nickname := mapper[user.CodeNickname]
	cookieValue := mapper[user.CodeCookie]
	username, err := redisutil.Get(cookieValue)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		return responseForbidden()
	}

	err = dbutil.UpdateNickname(nickname, username)
	if err != nil {
		log.Println(dbutil.ErrorUpdateProfile + " " + err.Error())
		return responseError(err)
	}
	user, err := dbutil.GetUser(username)
	setUserDataCache(user)

	return sendResponseBack(username, err)
}

func handleUpdateProfile(mapper map[string]string) response.Response{
	profile := mapper[user.CodeProfile]
	userIDFromCookie := mapper[user.CodeCookie]
	username, err := redisutil.Get(userIDFromCookie)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		return responseForbidden()
	}

	err = dbutil.UpdateProfile(profile, username)
	if err != nil {
		log.Println(dbutil.ErrorUpdateProfile + " " + err.Error())
		return responseError(err)
	}
	user, err := dbutil.GetUser(username)
	setUserDataCache(user)

	return sendResponseBack(username, err)
}

func handleRequest(requestID int32, mapper map[string]string) response.Response {
	if requestID == request.RequestLogin {
		return handleAuthentication(mapper)
	} else if requestID == request.RequestUserInfo {
		return handleInfo(mapper)
	} else if requestID == request.RequestUpdateNickname {
		return handleUpdateNickname(mapper)
	} else if requestID == request.RequestUpdateProfile {
		return handleUpdateProfile(mapper)
	}
	//should never reach this
	return response.Response{}
}

func main() {
	gob.Register(user.User{})
	gob.Register(http.Cookie{})
	gob.Register(response.Response{})

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

	grpcServer := grpc.NewServer()
	pb.RegisterUserDataServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
