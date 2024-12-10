package main

import (
	"log"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/command"
	"github.com/Caknoooo/go-gin-clean-starter/config"
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/Caknoooo/go-gin-clean-starter/middleware"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"github.com/Caknoooo/go-gin-clean-starter/routes"
	"github.com/Caknoooo/go-gin-clean-starter/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup Database Connection
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Command Line Argument Handler
	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	// JWT Service
	jwtService := service.NewJWTService()

	// User Implementation (Existing Setup)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtService)
	userController := controller.NewUserController(userService)

	// Mobil Implementation (New Setup)
	mobilRepository := repository.NewMobilRepository(db)
	mobilService := service.NewMobilService(mobilRepository, db)
	mobilController := controller.NewMobilController(mobilService)

	// Gin Server Setup
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// Routes
	routes.User(server, userController, jwtService)
	routes.MobilRoutes(server, mobilController)

	// Static Files
	server.Static("/assets", "./assets")

	// Port Setup
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	// Start Server
	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
