package api

import (
	"keycloak-app/services/handler"
	"log"

	"github.com/gin-gonic/gin"
	cfg "keycloak-app/config"
)

// StartServer inicializa o servidor e registra os handlers
func StartServer() {
	router := gin.Default()

	handler.SetupHandlers(router)

	log.Printf("✅ Server running on port %s", cfg.PORT)
	if err := router.Run(cfg.PORT); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
