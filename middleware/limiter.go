package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func GetLimiterMiddleware() interface{} {
	return limiter.New()
}
