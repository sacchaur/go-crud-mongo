package controllers

import (
	"crud_operation/configs"
	"crud_operation/libraries"
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
	Username string `json:"username"`
	jwt.StandardClaims
}

func (ctrl *authenticationController) Token(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Authenticate the user (you need to implement this function)
	status, err := ctrl.authLib.AuthenticateUser(username, password)
	if !status {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to authenticate user", "error": err.Error()})
	}

	expirationTime := time.Now().Add(configs.TokenExpiry)
	claims := &Claims{
		Username: username,
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

	c.Locals("username", claims.Username)
	return c.Next()
}

func (ctrl *authenticationController) Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Login page"})
}

func (ctrl *authenticationController) ResetPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Reset password page"})
}
