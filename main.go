package main

import (
	"crud_operation/configs"
	"crud_operation/repository"
	"crud_operation/routers"
	"crud_operation/stderrors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "crud_operation/docs"
)

// @title Go CRUD API with MongoDB
// @version 1.0.0
// @description This is a simple CRUD (Create, Read, Update, Delete) API written in Go, using the Fiber framework and MongoDB for storage.
// @contact.name Sachin Chaurasiya
// @contact.email chaurasia3011@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
// @produce json

// swagger:meta
// This is a placeholder comment.
// It will be ignored by the compiler, but it tells the swagger generation tool where to put the security definitions.

// swagger:securitySchemes
var ApiKeyAuth string

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: stderrors.Handler(),
	})

	//init env
	cfg, err := configs.NewApiConfig()
	if err != nil {
		log.Fatal(err)
	}

	// allow fiber to handle panics
	app.Use(recover.New())

	// setup storage connections
	err = repository.Init(cfg)
	if err != nil {
		log.Fatal("unable to connect to storage instances", err.Error())
	}

	// Register routes
	routers.AuthRoutes(app)

	routers.SetupRoutes(app, repository.StorageInstance)

	go startServer(cfg, app)
	gracefulShutdown(app)
}

// startServer listens to the configured port and starts the server
func startServer(cfg configs.ApiConfig, app *fiber.App) {
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal("error starting API server:", err)
	}
}

// gracefulShutdown handles the shutdown of the server when an interrupt or terminate signal is received
func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// attempt to gracefully shut down the server
	if err := app.Shutdown(); err != nil {
		log.Fatal("error shutting down server:", err)
	}
}
