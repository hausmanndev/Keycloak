package main

import (
	"context"
	"encoding/json"
	conf "keycloak-app/config"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func main() {
	conf.LoadConfig("./cmd/.env")
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     conf.CLIENT_ID,
		ClientSecret: conf.CLIENT_SECRET,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(conf.STATE), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != conf.STATE {
			http.Error(w, "Error: State does not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Error: It is not possible to exchange the token.", http.StatusInternalServerError)
			return
		}

		idToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "Error: It's not possible to retrieve the ID Token.", http.StatusInternalServerError)
			return
		}

		res := struct {
			OAuth2Token *oauth2.Token
			IDToken     string
		}{
			OAuth2Token: oauth2Token,
			IDToken:     idToken,
		}

		data, _ := json.MarshalIndent(res, "", "  ")
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
