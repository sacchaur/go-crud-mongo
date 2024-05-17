package routers

import (
	"crud_operation/controllers"
	"crud_operation/libraries"
	"crud_operation/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, storageInstance *mongo.Client) {
	// Initialize repositories
	userRepo := repository.NewUserRepository()
	userLibrary := libraries.NewUserService(userRepo)
	userController := controllers.NewUserHandler(userLibrary)

	// Route handlers
	app.Get("/users/:userid", userController.Get)
	app.Get("/users", userController.GetAll)
	app.Post("/users", userController.Add)
	app.Put("/users/:userid", userController.Update)
	app.Delete("/users/:userid", userController.Delete)

}
