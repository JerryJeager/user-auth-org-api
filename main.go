package main

import (
	"log"

	"github.com/JerryJeager/user-auth-org-api/cmd"
	"github.com/JerryJeager/user-auth-org-api/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	log.Print("Starting the user auth-organisation server...")
	cmd.ExecuteApiRoutes()
}
