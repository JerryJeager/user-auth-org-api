package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/JerryJeager/user-auth-org-api/middleware"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello from user-auth-org server"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
