package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rithvik89/auth/utils"
)

func Roothandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<a href="/login/github">LOGIN</a>`)
}

func GithubLoginHandler(rw http.ResponseWriter, r *http.Request) {

	// Create the dynamic redirect URL for login
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s",
		utils.GetGithubClientID(),
	)

	http.Redirect(rw, r, redirectURL, 301)
}

// Client side Middleware-Handler. (GithubCallbackHandler)

func GithubCallbackHandler(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		access_token := getGHAccessToken(code)

		user_data := getUserData(access_token)

		// Add this in request context for further Processing of request.

		ctx := context.WithValue(r.Context(), "data", user_data)

		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})

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

	type AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	ghResp := AccessTokenResponse{}

	respBody, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &ghResp)
	if err != nil {
		fmt.Printf("unable to unmarshal the response %s", err)
	}

	return ghResp.AccessToken

}
