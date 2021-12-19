package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetCORSMiddleware() interface{} {
	return cors.New(
		cors.Config{
			AllowOrigins:     "https://clump-fe.azurewebsites.net/",
			AllowCredentials: true,
			ExposeHeaders:    "*"},
	)
}
