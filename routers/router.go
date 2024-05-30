package routers

import (
	"crud_operation/controllers"
	"crud_operation/libraries"
	"crud_operation/middlewares"
	"crud_operation/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, storageInstance *mongo.Client) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Initialize repositories
	userRepo := repository.NewUserRepository()
	userLibrary := libraries.NewUserService(userRepo)
	userController := controllers.NewUserHandler(userLibrary)

	// add app to group v1
	v1 := app.Group("/v1")
	// Add v1 with middleware middlewares.JWTProtected()
	v1.Use(middlewares.JWTProtected())

	// Route handlers
	v1.Get("/users/:userid", userController.Get)
	v1.Get("/users", userController.GetAll)
	v1.Post("/users", userController.Add)
	v1.Put("/users/:userid", userController.Update)
	v1.Delete("/users/:userid", userController.Delete)

}

func AuthRoutes(app *fiber.App) {

	// Initialize repositories
	authRepo := repository.NewAuthenticationRepository()
	userLibrary := libraries.NewAuthenticationService(authRepo)
	oauthController := controllers.NewAuthenticationHandler(userLibrary)

	app.Post("/oauth/token", oauthController.Token)
	// app.Get("/protected", middlewares.JWTProtected(), func(c *fiber.Ctx) error {
	// 	username := c.Locals("username").(string)
	// 	return c.JSON(fiber.Map{"message": "Hello " + username})
	// })
	// app.Get("/oauth/callback", func(c *fiber.Ctx) error {
	// 	// Handle the OAuth callback here
	// 	return c.JSON(fiber.Map{"message": "OAuth callback received"})
	// })

	app.Post("/user/login", oauthController.Login)
	app.Post("/user/reset", oauthController.ResetPassword)

}
