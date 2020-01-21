package main

import "net/http"

func handleRouting() {
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/authenticate", handleAuthenticate)
	http.HandleFunc("/register/", handleRegister)
	http.HandleFunc("/registerAuth", handleRegisterAuth)
	http.HandleFunc("/info/", handleInfo)

	http.HandleFunc("/editnickname/", handleEditNickname)
	http.HandleFunc("/editprofile/", handleEditProfile)
	http.HandleFunc("/changenickname", handleChangeNickname)
	http.HandleFunc("/changeprofile", handleChangeProfile)
}