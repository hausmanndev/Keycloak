package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	cfg "keycloak-app/config"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

func Oauth2Config() *oauth2.Config {
	provider, err := oidc.NewProvider(ctx, cfg.KEYCLOAK_ISSUER)
	if err != nil {
		log.Fatal(err)
	}

	config := &oauth2.Config{
		ClientID:     cfg.CLIENT_ID,
		ClientSecret: cfg.CLIENT_SECRET,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost" + cfg.PORT + "/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	return config
}

func SetupHandlers(r *gin.Engine) {
	// Setup your route handlers here
	keycloak := r.Group("/keycloak")

	keycloak.GET("/", redirectToLogin())
	keycloak.GET("/auth/callback", authCallback())
}

// redirectToLogin redireciona o usuário para a tela de autenticação do Keycloak
func redirectToLogin() gin.HandlerFunc{
	return func(c *gin.Context) {
		config := Oauth2Config()
		authURL := config.AuthCodeURL(cfg.STATE)
		c.Redirect(http.StatusFound, authURL)
	}
}

// authCallback recebe o código do Keycloak e troca pelo token
func authCallback() gin.HandlerFunc{
	return func(c *gin.Context) {
		config := Oauth2Config()
		if c.Query("state") != cfg.STATE {
			c.JSON(http.StatusBadRequest, gin.H{"error": "state does not match"})
			return
		}

		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing authorization code"})
			return
		}

		oauth2Token, err := config.Exchange(ctx, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange token"})
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get ID token"})
			return
		}

		response := struct {
			OAuth2Token *oauth2.Token `json:"oauth2_token"`
			IDToken     string        `json:"id_token"`
		}{
			OAuth2Token: oauth2Token,
			IDToken:     rawIDToken,
		}

		data, _ := json.MarshalIndent(response, "", "  ")
		c.Data(http.StatusOK, "application/json", data)
	}
}
