package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"mohit.com/ecom-api/models"
)

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
