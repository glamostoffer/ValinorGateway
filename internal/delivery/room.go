package delivery

import (
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorGateway/pkg/consts"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateRoom(c *fiber.Ctx) error {
	//user, ok := c.Locals(consts.UserLocalsKey).(*model.UserLocals)
	//if !ok {
	//	return c.SendStatus(fiber.StatusUnauthorized)
	//}

	request := model.CreateRoomRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	roomID, err := h.uc.CreateRoom(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(roomID)
}

func (h *Handler) GetClientRooms(c *fiber.Ctx) error {
	user, ok := c.Locals(consts.UserLocalsKey).(*model.UserLocals)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	rooms, err := h.uc.GetListOfRooms(c.Context(), user.UserID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(model.GetListOfRoomsResponse{Rooms: rooms})
}

func (h *Handler) AddClientToRoom(c *fiber.Ctx) error {
	user, ok := c.Locals(consts.UserLocalsKey).(*model.UserLocals)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	request := model.AddClientToRoomRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	if request.OwnerID != user.UserID {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err := h.uc.AddClientToRoom(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) RemoveClientFromRoom(c *fiber.Ctx) error {
	user, ok := c.Locals(consts.UserLocalsKey).(*model.UserLocals)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	request := model.RemoveClientFromRoomRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	if request.OwnerID != user.UserID {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err := h.uc.RemoveClientFromRoom(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
