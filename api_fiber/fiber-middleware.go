package api_fiber

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		diff := time.Since(start)
		c.Append("X-Response-Time", fmt.Sprintf("%dms", diff.Milliseconds()))

		return err
	}
}
