package controllers

import (
	"crud_operation/configs"
	"crud_operation/dto"
	"crud_operation/responses"
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

// User Status   godoc
// @Summary      Get user
// @Description  Retrieve current user by id
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserStatus
// @Router       /api/v1/user/status [get]
// @Security     ApiKeyAuth
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
			Data:    &fiber.Map{"data": "User not found."},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    &fiber.Map{"data": user},
	})
}

func (userController *userController) Add(c *fiber.Ctx) error {
	log.Println("Add user in controller")
	// Read the request body
	// Unmarshal the request body into the User struct
	var user dto.User
	json.Unmarshal(c.Body(), &user)

	response, err := userController.userLib.Add(c.UserContext(), &user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    &fiber.Map{"data": response},
	})
}

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
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    &fiber.Map{"data": response},
	})
}

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
			Data:    &fiber.Map{"data": "User delete successfully."},
		})
	}

	// If the user is not found, respond with an error
	return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
		Status:  http.StatusNotFound,
		Message: dto.Error,
		Data:    &fiber.Map{"data": "User not found."},
	})
}

func (userController *userController) GetAll(c *fiber.Ctx) error {
	log.Println("Get all user in controller1")
	users, err := userController.userLib.GetAll(c.UserContext())
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: dto.Error,
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: dto.Success,
		Data:    &fiber.Map{"data": users},
	})
}
