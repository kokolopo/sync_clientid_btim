package handlers

import (
	"sync_btim/service"

	"github.com/gofiber/fiber/v2"
)

type H map[string]interface{}

type TablenameHandler struct {
	tablenameService service.IService
}

func NewTablenameHandler(TablenameService service.IService) *TablenameHandler {
	return &TablenameHandler{TablenameService}
}

func (h *TablenameHandler) SyncClientIDBTIM(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 10)
	email := c.Query("email")

	res, err := h.tablenameService.UpdateClientID(email, limit, offset)
	if err != nil {
		return c.JSON(H{
			"error": err.Error(),
		})
	}

	return c.JSON(H{
		"is_migrate": res})
}
