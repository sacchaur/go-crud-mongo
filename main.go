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
)

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
