package main

import (
	"github.com/go-chi/chi/v5"
	auth "github.com/rithvik89/auth/pkg/auth"
)

func initHandler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/auth/userpass", UserPassAuthHandler(r))
	r.Route("/login", GithubHandler(r))

	return r
}

func UserPassAuthHandler(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", handleLogin)
		r.Post("/signup", handleSignup)
		r.Delete("/logout", handleLogout)
	}
}

func GithubHandler(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", auth.Roothandler)
		r.Get("/github", auth.GithubLoginHandler)

		// make sure this acts as middleware enabled route.
		r.Get("/github/callback", auth.GithubCallbackHandler(handleGithubUserdata))

	}
}
