package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/JerryJeager/user-auth-org-api/docs"
	"github.com/JerryJeager/user-auth-org-api/manualwire"
	"github.com/JerryJeager/user-auth-org-api/middleware"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()
var orgController = manualwire.GetOrgController()

func ExecuteApiRoutes() {

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello from user-auth-org server"})
	})
	api := router.Group("/api")

	api.GET("/info/openapi.yaml", func(c *gin.Context) {
		c.String(200, docs.OpenApiDocs())
	})

	auth := router.Group("/auth")
	auth.POST("/register", userController.CreateUser)
	auth.POST("/login", userController.LoginUser)

	api.GET("/users/:id", middleware.JwtAuthMiddleware(), userController.GetUser)

	organisations := api.Group("/organisations")
	organisations.POST("", middleware.JwtAuthMiddleware(), orgController.CreateOrganisation)
	organisations.POST("/:orgId/users", orgController.CreateOrgMember)
	organisations.GET("/:orgId", orgController.GetOrganisation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
