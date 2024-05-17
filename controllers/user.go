package controllers

import (
	"crud_operation/configs"
	"crud_operation/dto"
	"crud_operation/responses"
	"crud_operation/stderrors"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"crud_operation/libraries"

	"github.com/gofiber/fiber/v2"
)

// UserController is the interface for the user controller
type UserController interface {
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type userController struct {
	config  configs.ApiConfig
	userLib libraries.UserService
}

func NewUserHandler(userLib libraries.UserService) UserController {
	return &userController{
		config:  configs.GetConfig(),
		userLib: userLib,
	}
}

// Get retrieves a user by ID.
// @Summary Get a user by ID
// @Description Retrieves a user by the given ID.
// @Tags Users
// @Param userid path int true "User ID"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {object} responses.UserResponse
// @Router /users/{userid} [get]
func (userController *userController) Get(c *fiber.Ctx) error {
	// Get the user id from the param
	log.Println("Get user in controller")
	userId := c.Params("userid", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return c.SendStatus(400)
	}
	user, err := userController.userLib.Get(c.UserContext(), userIdInt)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Error:   stderrors.ErrNotFound(c.UserContext(), "User not found"),
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    user,
	})
}

// Add creates a new user.
// @Summary Create a new user
// @Description Creates a new user with the provided details.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Router /users [post]
func (userController *userController) Add(c *fiber.Ctx) error {
	log.Println("Add user in controller")
	// Read the request body
	// Unmarshal the request body into the User struct
	var user dto.User
	json.Unmarshal(c.Body(), &user)

	// Check if the user already exists
	existingUser, err := userController.userLib.Get(c.UserContext(), user.UserId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Error:   err,
		})
	}

	if existingUser != nil {
		return c.Status(http.StatusConflict).JSON(responses.UserResponse{
			Status:  http.StatusConflict,
			Message: dto.Error,
			Error:   stderrors.ErrConflict(c.UserContext(), "User already exists."),
		})
	}

	response, err := userController.userLib.Add(c.UserContext(), &user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    response,
	})
}

// Update updates an existing user.
// @Summary Update an existing user
// @Description Updates an existing user with the provided details.
// @Tags Users
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Param user body dto.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {object} responses.UserResponse
// @Router /users/{userid} [put]
func (userController *userController) Update(c *fiber.Ctx) error {
	userId := c.Params("userid", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return c.SendStatus(400)
	}

	// Read the request body
	var updatedUser dto.User
	// Unmarshal the request body into the User struct
	err = json.Unmarshal(c.Body(), &updatedUser)
	if err != nil {
		// If the request body is not valid, respond with an error
		return err
	}

	response, err := userController.userLib.Update(c.UserContext(), userIdInt, &updatedUser)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Error:   err,
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    response,
	})
}

// Delete deletes a user by ID.
// @Summary Delete a user by ID
// @Description Deletes a user by the given ID.
// @Tags Users
// @Param userid path int true "User ID"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {object} responses.UserResponse
// @Router /users/{userid} [delete]
func (userController *userController) Delete(c *fiber.Ctx) error {
	userId := c.Params("userid", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return c.SendStatus(400)
	}

	status, err := userController.userLib.Delete(c.UserContext(), userIdInt)
	if err != nil {
		return c.SendStatus(404)
	}

	if status {
		return c.Status(http.StatusOK).JSON(responses.UserResponse{
			Status:  http.StatusOK,
			Message: dto.Success,
		})
	}

	// If the user is not found, respond with an error
	return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
		Status:  http.StatusNotFound,
		Message: dto.Error,
		Error:   stderrors.ErrNotFound(c.UserContext(), "User not found"),
	})
}

// GetAll retrieves all users.
// @Summary Get all users
// @Description Retrieves all users.
// @Tags Users
// @Success 200 {object} responses.UsersResponse
// @Failure 404 {object} responses.UsersResponse
// @Router /users [get]
func (userController *userController) GetAll(c *fiber.Ctx) error {
	log.Println("Get all user in controller")
	users, err := userController.userLib.GetAll(c.UserContext())
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Error:   err,
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UsersResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    users,
	})
}
