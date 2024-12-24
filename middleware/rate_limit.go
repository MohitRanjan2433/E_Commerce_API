package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiterMiddleware(maxRequests int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max: maxRequests,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwaded-for", c.IP())
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html")
		},
		Storage: memory.New(),
	})
}
