package httpserver

import (
	"context"
	"errors"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dickynovanto1103/User-Management-System/container"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"
	"github.com/dickynovanto1103/User-Management-System/internal/service/authentication"
	"github.com/dickynovanto1103/User-Management-System/internal/service/cookie"
	"github.com/dickynovanto1103/User-Management-System/internal/service/fileuploader"
	pb "github.com/dickynovanto1103/User-Management-System/proto"
)

var templates = template.Must(template.ParseGlob("internal/service/httpServer/templates/*"))

func handleHome(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Println("error executing template home", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	var conn net.Conn
	var status string
	resp, err := sendLoginInformation(r, conn)
	if err != nil {
		log.Println("sending login information failed:", err)
		return
	}
	data := resp.GetMapper()
	status = data[model.ResponseCode]

	if status == authentication.ErrorNotAuthenticated || status == dbsql.ErrorGetPassword {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sessionID := status
	cookie.CreateCookie(w, sessionID, 5*time.Hour)
	http.Redirect(w, r, "/info", http.StatusPermanentRedirect)

	err = templates.ExecuteTemplate(w, "info.html", nil)
	if err != nil {
		log.Println("error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handleRegisterAuth(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	nickname := r.FormValue("nickname")
	password := r.FormValue("password")
	log.Println(username, nickname, password)
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
	nickname := r.FormValue(model.CodeNickname)

	var mapper = make(map[string]string)
	cookie, err := cookie.GetCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	mapper[model.CodeNickname] = nickname
	mapper[model.CodeCookie] = cookie.Value

	resp, err := sendRequest(model.RequestUpdateNickname, mapper, container.Client)
	if err != nil {
		return
	}

	err = readResponse(w, r, resp)
	if err != nil {
		return
	}
}

func handleChangeProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(model.CodeSessionID)
	if err != nil {
		log.Println("error getting cookie", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	filename, err := fileuploader.UploadFile(r, model.CodeProfile, cookie.Value)

	var mapper = make(map[string]string)
	mapper[model.CodeProfile] = filename
	mapper[model.CodeCookie] = cookie.Value

	resp, err := sendRequest(model.RequestUpdateProfile, mapper, container.Client)
	if err != nil {
		return
	}
	err = readResponse(w, r, resp)
}

func sendRequest(requestID int, mapper map[string]string, client pb.UserDataServiceClient) (*pb.Response, error) {
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

func convertRequestPBToRequestStructure(resp *pb.Response) model.Response {
	response := model.Response{
		ResponseID: resp.GetResponseID(),
		Data:       resp.GetMapper(),
	}
	return response
}

func readResponse(w http.ResponseWriter, r *http.Request, responsePB *pb.Response) error {
	resp := convertRequestPBToRequestStructure(responsePB)
	if resp.ResponseID == model.ResponseForbidden || resp.ResponseID == model.ResponseError {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return nil
	}

	userData := resp.Data[model.CodeUser]
	http.Redirect(w, r, "/info", http.StatusSeeOther)
	templates.ExecuteTemplate(w, "info.html", userData)

	return nil
}

func sendLoginInformation(r *http.Request, conn net.Conn) (*pb.Response, error) {
	username := r.FormValue(model.CodeUsername)
	password := r.FormValue(model.CodePassword)

	var mapper = make(map[string]string)
	mapper[model.CodeUsername] = username
	mapper[model.CodePassword] = password
	resp, err := sendRequest(model.RequestLogin, mapper, container.Client)
	return resp, err
}

func getUserDataFromTCPServer(r *http.Request) (model.User, error) {
	cookie, err := cookie.GetCookie(r)
	if err != nil {
		log.Println("error retrieving cookie: ", err)
		return model.User{}, err
	}

	var mapper = make(map[string]string)
	mapper[model.CodeCookie] = cookie.Value
	resp, err := sendRequest(model.RequestUserInfo, mapper, container.Client)

	if err != nil {
		log.Println("error in receiving response: ", err)
		return model.User{}, errors.New(model.ResponseKeyForbidden)
	}

	if resp.ResponseID == model.ResponseForbidden {
		return model.User{}, errors.New(model.ResponseKeyForbidden)
	}

	userData := getUserFromDataInResponse(resp.GetMapper())
	return userData, nil
}

func getUserFromDataInResponse(mapper map[string]string) model.User {
	var userResult model.User
	userResult.Username = mapper[model.CodeUsername]
	userResult.Nickname = mapper[model.CodeNickname]
	userResult.ProfileURL = mapper[model.CodeProfile]

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
