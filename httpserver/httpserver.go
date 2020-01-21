package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dickynovanto1103/User-Management-System/internal/authentication"
	"github.com/dickynovanto1103/User-Management-System/internal/connection"
	"github.com/dickynovanto1103/User-Management-System/internal/cookie"
	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"
	"github.com/dickynovanto1103/User-Management-System/internal/fileuploader"
	"github.com/dickynovanto1103/User-Management-System/internal/request"
	"github.com/dickynovanto1103/User-Management-System/internal/response"

	"github.com/dickynovanto1103/User-Management-System/internal/user"
	pb "github.com/dickynovanto1103/User-Management-System/proto"
)

const CodeForbidden = "Forbidden"

var connPool = &connection.ConnPool{}
var templates = template.Must(template.ParseGlob("templates/*"))
var client pb.UserDataServiceClient

func sendRequest(requestID int, mapper map[string]string) (*pb.Response, error) {
	request := &pb.Request{
		RequestID:            int32(requestID),
		Mapper:               mapper,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	response, err := client.SendRequest(context.Background(), request)
	return response, err
}

func convertRequestPBToRequestStructure(resp *pb.Response) response.Response {
	response := response.Response{
		ResponseID: resp.GetResponseID(),
		Data:       resp.GetMapper(),
	}
	return response
}

func readResponse(w http.ResponseWriter, r *http.Request, responsePB *pb.Response) error {
	resp := convertRequestPBToRequestStructure(responsePB)
	if resp.ResponseID == response.ResponseForbidden || resp.ResponseID == response.ResponseError {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		userData := resp.Data[user.CodeUser]
		http.Redirect(w, r, "/info", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "info.html", userData)
	}
	return nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println("error executing template login", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		log.Println("error in executing template register", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sendLoginInformation(r *http.Request, conn net.Conn) (*pb.Response, error) {
	username := r.FormValue(user.CodeUsername)
	password := r.FormValue(user.CodePassword)

	var mapper = make(map[string]string)
	mapper[user.CodeUsername] = username
	mapper[user.CodePassword] = password
	resp, err := sendRequest(request.RequestLogin, mapper)
	return resp, err
}

func handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	var conn net.Conn
	var status string
	resp, err := sendLoginInformation(r, conn)
	if err != nil {
		log.Println("sending login information failed:", err)
		return
	}
	data := resp.GetMapper()
	status = data[response.ResponseCode]

	if status == authentication.ErrorNotAuthenticated || status == dbutil.ErrorGetPassword {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		sessionID := status
		cookie.CreateCookie(w, sessionID, 5*time.Hour)
		http.Redirect(w, r, "/info", http.StatusPermanentRedirect)
		err := templates.ExecuteTemplate(w, "info.html", nil)
		if err != nil {
			log.Println("error executing template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleRegisterAuth(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	nickname := r.FormValue("nickname")
	password := r.FormValue("password")
	log.Println(username, nickname, password)
}

func getUserDataFromTCPServer(r *http.Request) (user.User, error) {
	cookie, err := cookie.GetCookie(r)
	if err != nil {
		log.Println("error retrieving cookie: ", err)
		return user.User{}, err
	}
	var mapper = make(map[string]string)
	mapper[user.CodeCookie] = cookie.Value
	resp, err := sendRequest(request.RequestUserInfo, mapper)

	if err != nil {
		log.Println("error in receiving response: ", err)
		return user.User{}, errors.New(response.ResponseKeyForbidden)
	}

	if resp.ResponseID == response.ResponseForbidden {
		return user.User{}, errors.New(CodeForbidden)
	} else {
		userData := getUserFromDataInResponse(resp.GetMapper())
		return userData, nil
	}
}

func getUserFromDataInResponse(mapper map[string]string) user.User {
	var userResult user.User
	userResult.Username = mapper[user.CodeUsername]
	userResult.Nickname = mapper[user.CodeNickname]
	userResult.ProfileURL = mapper[user.CodeProfile]

	return userResult
}

func showPage(w http.ResponseWriter, r *http.Request, pageName string) {
	userData, err := getUserDataFromTCPServer(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	templates.ExecuteTemplate(w, pageName+".html", userData)
}

func handleEditNickname(w http.ResponseWriter, r *http.Request) {
	showPage(w, r, "editnickname")
}

func handleEditProfile(w http.ResponseWriter, r *http.Request) {
	showPage(w, r, "editprofile")
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	showPage(w, r, "info")
}

func handleChangeNickname(w http.ResponseWriter, r *http.Request) {
	nickname := r.FormValue(user.CodeNickname)

	var mapper = make(map[string]string)
	cookie, err := cookie.GetCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	mapper[user.CodeNickname] = nickname
	mapper[user.CodeCookie] = cookie.Value

	resp, err := sendRequest(request.RequestUpdateNickname, mapper)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
	err = readResponse(w, r, resp)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
}

func handleChangeProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(user.CodeSessionID)
	if err != nil {
		log.Println("error getting cookie", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	filename, err := fileuploader.UploadFile(r, user.CodeProfile, cookie.Value)

	var mapper = make(map[string]string)
	mapper[user.CodeProfile] = filename
	mapper[user.CodeCookie] = cookie.Value

	resp, err := sendRequest(request.RequestUpdateProfile, mapper)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
	err = readResponse(w, r, resp)
}

func main() {
	handleRouting()

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error dialing grpc: ", err)
	}
	defer conn.Close()

	client = pb.NewUserDataServiceClient(conn)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
