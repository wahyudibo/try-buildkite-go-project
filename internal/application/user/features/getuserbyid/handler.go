package getuserbyid

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	userdao "github.com/wahyudibo/try-buildkite-go-project/internal/application/user/dao"
)

// Handler handles incoming http connection from getuserbyid routes.
type Handler struct {
	userDAO *userdao.DAO
}

func New(userDAO *userdao.DAO) *Handler {
	return &Handler{
		userDAO: userDAO,
	}
}

func (h *Handler) Handler(c *fiber.Ctx) error {
	userIDParam := c.Params("userId")
	if userIDParam == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	user, err := h.userDAO.GetByID(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"user": user,
	})
}
