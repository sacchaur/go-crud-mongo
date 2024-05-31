package configs

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWTSecret   = []byte(os.Getenv("JWT_SECRET"))
	TokenExpiry = time.Hour * 24
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
