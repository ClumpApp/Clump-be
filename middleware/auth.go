package middleware

import (
	"time"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

const key = "ASE"

func GetJWTMiddleware() interface{} {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(key),
	})
}

func GetToken(name string) string {
	claims := jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}

	return t
}

func GetClaim(token interface{}) string {
	user := token.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["name"].(string)
}
