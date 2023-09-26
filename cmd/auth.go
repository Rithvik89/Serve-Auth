package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rithvik89/auth/utils"
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

func handleGithubLogin(rw http.ResponseWriter, r *http.Request) {
	// client_secret and client_id from secrets

}

func roothandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<a href="/login/github">LOGIN</a>`)
}

func githubLoginHandler(rw http.ResponseWriter, r *http.Request) {

	// Create the dynamic redirect URL for login
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s",
		utils.GetGithubClientID(),
	)

	http.Redirect(rw, r, redirectURL, 301)
}

func githubCallbackHandler(rw http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	access_token := getGHAccessToken(code)

	user_data := getUserData(access_token)

	fmt.Fprintf(rw, user_data)
}

func getUserData(token string) string {
	rq, _ := http.NewRequest("GET", "https://api.github.com/user", nil)

	token_ := fmt.Sprintf("Bearer %s", token)

	rq.Header.Set("Authorization", token_)

	res, err := http.DefaultClient.Do(rq)

	if err != nil {
		fmt.Printf("Error occured %s", err)
	}

	respBody, _ := io.ReadAll(res.Body)
	return string(respBody)

}

func getGHAccessToken(code string) string {

	type ATReqDetails struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}

	reqBody := ATReqDetails{
		ClientId:     utils.GetGithubClientID(),
		ClientSecret: utils.GetGHClientSecret(),
		Code:         code,
	}

	reqBodyEncode, err := json.Marshal(reqBody)

	if err != nil {
		fmt.Printf("unable to marshal the request %s", err)
	}

	reqBodyBuff := bytes.NewBuffer(reqBodyEncode)

	rq, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", reqBodyBuff)

	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("Accept", "application/json")

	resp, _ := http.DefaultClient.Do(rq)

	respBody, _ := io.ReadAll(resp.Body)

	type AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	ghResp := AccessTokenResponse{}

	err = json.Unmarshal(respBody, &ghResp)
	if err != nil {
		fmt.Printf("unable to unmarshal the response %s", err)
	}

	return ghResp.AccessToken

}
