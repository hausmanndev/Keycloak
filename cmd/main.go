package main

import (
	cfg "keycloak-app/config"
	"keycloak-app/services/api"
)

func main() {
	// Carrega variáveis de ambiente
	cfg.LoadConfig("./.env")

	// Inicializa o servidor com Gin
	api.StartServer()
}
