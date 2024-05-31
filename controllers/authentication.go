package controllers

import (
	"crud_operation/configs"
	"crud_operation/dto"
	"crud_operation/libraries"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationController interface {
	Token(c *fiber.Ctx) error
	ValidateToken(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
}

type authenticationController struct {
	config  configs.ApiConfig
	authLib libraries.AuthenticationService
}

func NewAuthenticationHandler(authLib libraries.AuthenticationService) AuthenticationController {
	return &authenticationController{
		config:  configs.GetConfig(),
		authLib: authLib,
	}
}

type Claims struct {
	ClientId string `json:"client_id"`
	jwt.StandardClaims
}

// Create token
// @Summary Create token
// @Description Create token
// @Tags Authentication
// @Produce json
// @Accept application/x-www-form-urlencoded
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} responses.TokenResponse
// @Failure 401 {object} responses.ErrorResponse "Invalid credentials"
// @Failure 500 {object} responses.ErrorResponse "Failed to create token"
// @Router /oauth/token [post]
func (ctrl *authenticationController) Token(c *fiber.Ctx) error {
	clientId := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	// Authenticate the user (you need to implement this function)
	status, err := ctrl.authLib.AuthenticateToken(clientId, clientSecret)
	if !status {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to authenticate user", "error": err.Error()})
	}

	expirationTime := time.Now().Add(configs.TokenExpiry)
	claims := &Claims{
		ClientId: clientId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(configs.JWTSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func (ctrl *authenticationController) ValidateToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return configs.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("clinet_id", claims.ClientId)
	return c.Next()
}

// Login
// @Summary Login
// @Description Login
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Router /user/login [post]
func (ctrl *authenticationController) Login(c *fiber.Ctx) error {
	// fetch the username and password from the request body
	var user dto.User
	json.Unmarshal(c.Body(), &user)
	status, err := ctrl.authLib.Login(user.Username, user.Password)
	if !status {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to authenticate user", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Login successful."})
}

// ResetPassword
// @Summary ResetPassword
// @Description ResetPassword
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse "Failed to reset password"
// @Router /user/reset [post]
func (ctrl *authenticationController) ResetPassword(c *fiber.Ctx) error {
	var user dto.User
	json.Unmarshal(c.Body(), &user)
	status, err := ctrl.authLib.ResetPassword(user.Username, user.Password)

	if err != nil || !status {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to reset password"})
	}

	return c.JSON(fiber.Map{"message": "Password reset successful."})
}
