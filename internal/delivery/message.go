package delivery

import (
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetMessagesFromRoom(c *fiber.Ctx) error {
	request := model.GetMessagesFromRoomRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	messages, err := h.uc.GetMessagesFromRoom(c.Context(), request.RoomID)
	if err != nil {
		return err
	}

	resp := model.GetMessagesFromRoomResponse{
		Messages: messages,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
