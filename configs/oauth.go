package configs

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWTSecret         = []byte(os.Getenv("JWT_SECRET"))
	OauthClientID     = os.Getenv("OAUTH_CLIENT_ID")
	OauthClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	TokenExpiry       = time.Hour * 24
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
