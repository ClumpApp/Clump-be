package middleware

import (
	"time"

	"github.com/clumpapp/clump-be/utility"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

const (
	group       = "group"
	user        = "user"
	ContextKey  = "account"
	tokenLookup = "header:Authorization,query:token"
)

func GetJWTMiddleware() interface{} {
	return jwtware.New(jwtware.Config{
		ContextKey:  ContextKey,
		TokenLookup: tokenLookup,
		SigningKey:  []byte(utility.GetConfig().GetJWTKey()),
	})
}

func CreateToken(userid, groupid uint) string {
	claims := jwt.MapClaims{
		user:  userid,
		group: groupid,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(utility.GetConfig().GetJWTKey()))
	if err != nil {
		return err.Error()
	}

	return t
}

func GetGroupID(token interface{}) float64 {
	jwtToken := token.(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims[group].(float64)
}

func GetUserID(token interface{}) float64 {
	jwtToken := token.(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims[user].(float64)
}
