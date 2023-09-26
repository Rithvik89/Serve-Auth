package main

import (
	"github.com/go-chi/chi/v5"
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
		r.Get("/", roothandler)
		r.Get("/github", githubLoginHandler)
		r.Get("/github/callback", githubCallbackHandler)

	}
}
