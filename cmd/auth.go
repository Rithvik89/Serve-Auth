package main

import (
	"net/http"
)

type LoginDetails struct {
	MailId   string `json:"mail_id"`
	Password string `json:"password"`
}

func handleLogin(rw http.ResponseWriter, r *http.Request) {
	var incBody LoginDetails
	err := getBody(r, &incBody)

	if err != nil {

	}

	if incBody.Password == "abcde" {
		sendRes(rw, 200, "Login Successful", nil)
	}
}

func handleSignup(rw http.ResponseWriter, r *http.Request) {

}

func handleLogout(rw http.ResponseWriter, r *http.Request) {

}

func handleGithubUserdata(rw http.ResponseWriter, r *http.Request) {

	user_data := r.Context().Value("data")

	sendRes(rw, 200, "GH user data retreived", user_data)
}
