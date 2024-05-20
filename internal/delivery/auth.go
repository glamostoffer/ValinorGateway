package delivery

import (
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorGateway/pkg/consts"
	"github.com/gofiber/fiber/v2"
	"time"
)

// ==================== ADMIN ==================== //

func (h *Handler) AdminSignUp(c *fiber.Ctx) error {
	request := model.AdminSignUpRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	err := h.uc.AdminSignUp(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) BanUser(c *fiber.Ctx) error {
	request := model.BanUserRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	err := h.uc.BanUser(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) CreateInviteToken(c *fiber.Ctx) error {
	request := model.CreateInviteTokenRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	response, err := h.uc.CreateInviteToken(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handler) GetListOfUsers(c *fiber.Ctx) error {
	request := model.GetListOfUsersRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	response, err := h.uc.GetListOfUsers(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// ==================== USER ==================== //

func (h *Handler) SignUp(c *fiber.Ctx) error {
	request := model.SignUpRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	err := h.uc.SignUp(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) GetClientDetails(c *fiber.Ctx) error {
	user, ok := c.Locals(consts.UserLocalsKey).(*model.UserLocals)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	request := model.GetClientDetailsRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	if user.UserID != request.ClientID && user.Role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	response, err := h.uc.GetClientDetails(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handler) UpdateClientDetails(c *fiber.Ctx) error {
	user, ok := c.Locals(consts.UserLocalsKey).(*model.UserLocals)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	request := model.UpdateClientDetailsRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	if user.UserID != request.ClientID { // && user.Role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if request.Login == nil && request.Password == nil {
		return c.SendStatus(fiber.StatusOK)
	}

	err := h.uc.UpdateClientDetails(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	request := model.SignInRequest{}
	if err := h.validateRequest(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.IncorrectParamsCause{FailCause: err.Error()})
	}

	response, err := h.uc.SignIn(c.Context(), request)
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:    "accessToken",
		Value:   response.Token,
		Path:    "",
		Domain:  "",
		MaxAge:  0,
		Expires: time.Now().Add(h.cfg.AccessTokenTTL),

		// Пока похуй, но поменять будто надо
		Secure:      false,
		HTTPOnly:    false,
		SameSite:    "",
		SessionOnly: false,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(response)
}
