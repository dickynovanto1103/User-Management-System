package main

import (
	"bufio"
	"encoding/gob"
	"errors"
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
)

const CodeForbidden = "Forbidden"

var connPool = &connection.ConnPool{}
var templates = template.Must(template.ParseGlob("templates/*"))

func sendRequest(conn net.Conn, encoder *gob.Encoder, requestID int, mapper map[string]interface{}) error {
	req := request.Request{RequestID: requestID, Data: mapper}
	err := encoder.Encode(req)
	if err != nil {
		log.Println("encoding error here: ", err)
		return err
	}
	return nil
}

func readResponse(w http.ResponseWriter, r *http.Request, conn net.Conn) error {
	var resp response.Response
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&resp)
	if err != nil {
		log.Println("error in decoding: ", err)
		return err
	}

	if resp.ResponseID == response.ResponseForbidden {
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

func sendLoginInformation(r *http.Request, conn net.Conn) error {
	username := r.FormValue(user.CodeUsername)
	password := r.FormValue(user.CodePassword)
	encoder := gob.NewEncoder(conn)

	var mapper = make(map[string]interface{})
	mapper[user.CodeUsername] = username
	mapper[user.CodePassword] = password
	err := sendRequest(conn, encoder, request.RequestLogin, mapper)
	return err
}

func handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	var conn net.Conn
	var status string
	finallySuccess := false
	for i := 0; i < request.MaxTries; i++ {
		conn = connPool.Get()

		err := sendLoginInformation(r, conn)
		if err != nil {
			log.Println("sending login information failed:", err)
			return
		}
		status, err = bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Println("error", err)
			connPool.CreateNewConnection()
			conn = connPool.Get()
		} else {
			finallySuccess = true
			break
		}
	}
	if !finallySuccess {
		log.Println("error not finallySuccess")
		http.Error(w, "error", http.StatusInternalServerError)
		connPool.CreateNewConnection()
		return
	}

	defer connPool.Put(conn)

	status = status[:len(status)-1] //to ignore last \n
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

func getUserDataFromTCPServer(w http.ResponseWriter, r *http.Request, conn net.Conn) (user.User, error) {
	cookie, err := cookie.GetCookie(r)
	if err != nil {
		log.Println("error retrieving cookie: ", err)
		return user.User{}, err
	}
	encoder := gob.NewEncoder(conn)
	var mapper = make(map[string]interface{})
	mapper[user.CodeCookie] = cookie
	sendRequest(conn, encoder, request.RequestUserInfo, mapper)

	var resp response.Response
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&resp)
	if err != nil {
		log.Println("error in decoding: ", err)
		return user.User{}, errors.New(response.ResponseKeyForbidden)
	}

	if resp.ResponseID == response.ResponseForbidden {
		return user.User{}, errors.New(CodeForbidden)
	} else {
		userData := resp.Data[user.CodeUser]
		return userData.(user.User), nil
	}
}

func showPage(w http.ResponseWriter, r *http.Request, pageName string) {
	conn := connPool.Get()
	defer connPool.Put(conn)
	userData, err := getUserDataFromTCPServer(w, r, conn)
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
	conn := connPool.Get()

	nickname := r.FormValue(user.CodeNickname)

	encoder := gob.NewEncoder(conn)

	var mapper = make(map[string]interface{})
	cookie, err := cookie.GetCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	mapper[user.CodeNickname] = nickname
	mapper[user.CodeCookie] = cookie

	err = sendRequest(conn, encoder, request.RequestUpdateNickname, mapper)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
	err = readResponse(w, r, conn)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
	connPool.Put(conn)
}

func handleChangeProfile(w http.ResponseWriter, r *http.Request) {
	conn := connPool.Get()

	cookie, err := r.Cookie(user.CodeSessionID)
	if err != nil {
		log.Println("error getting cookie", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	filename, err := fileuploader.UploadFile(r, user.CodeProfile, cookie.Value)
	encoder := gob.NewEncoder(conn)

	var mapper = make(map[string]interface{})
	mapper[user.CodeProfile] = filename
	mapper[user.CodeCookie] = cookie

	err = sendRequest(conn, encoder, request.RequestUpdateProfile, mapper)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
	err = readResponse(w, r, conn)
	if err != nil {
		connPool.CreateNewConnection()
		return
	}
	connPool.Put(conn)
}

func main() {
	gob.Register(user.User{})
	gob.Register(request.Request{})
	gob.Register(http.Cookie{})

	handleRouting()

	log.Println("start generating connections")
	err := connPool.CreatePool(connection.MaxConnections)
	if err != nil {
		log.Println("error creating connection pool", err)
	}
	log.Println("done generating connections")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
