package main

import (
	"fmt"

	"github.com/sachinzzzzz/foo/config"
	"github.com/sachinzzzzz/foo/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	host, user, password, dbname := config.GetDBConfig()
	db.ConnectDB(host, user, password, dbname)

	// Init router
	r := gin.Default()

	// // Register routes
	// routes.RegisterRoutes(r)

	// Start server
	port := config.GetPort()
	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
