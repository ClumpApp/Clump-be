package middleware

import (
	"time"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

const (
	key         = "ASE"
	Group       = "group"
	User        = "user"
	TokenSource = "cookie"
	TokenName   = "token"
	ContextKey  = "account"
)

func GetJWTMiddleware() interface{} {
	return jwtware.New(jwtware.Config{
		ContextKey:  ContextKey,
		SigningKey:  []byte(key),
		TokenLookup: TokenSource + ":" + TokenName,
	})
}

func CreateToken(user, group uint) string {
	claims := jwt.MapClaims{
		User:  user,
		Group: group,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}

	return t
}

func GetGroupID(token interface{}) float64 {
	user := token.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[Group].(float64)
}

func GetUserID(token interface{}) float64 {
	user := token.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[User].(float64)
}
