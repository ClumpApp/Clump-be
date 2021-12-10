package middleware

import (
	"time"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

const key = "ASE"
const Group = "group"
const User = "user"

func GetJWTMiddleware() interface{} {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(key),
	})
}

func GetToken(user, group string) string {
	claims := jwt.MapClaims{
		User: user,
		Group: group,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}

	return t
}

func GetGroupID(token interface{}) string {
	user := token.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[Group].(string)
}

func GetUserID(token interface{}) string {
	user := token.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[User].(string)
}
