package api_fiber

import (
	"github.com/VDiPaola/api-network-server/models"
	"github.com/VDiPaola/api-network-server/nodes"
	"github.com/gofiber/fiber/v2"
)

func NodeHealthCheck(c *fiber.Ctx) error {
	//endpoint to add/update node
	var node models.Node
	if err := c.BodyParser(&node); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	nodes.HealthCheck(node)

	return c.Status(200).SendString("")
}
