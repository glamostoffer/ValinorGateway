package delivery

import (
	"github.com/glamostoffer/ValinorGateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(route fiber.Router, mw middleware.Middleware, h *Handler) {
	route.Post("sign-up", h.SignUp)
	route.Post("sign-in", h.SignIn)

	// GetClientDetails и UpdateClientDetails - подумать, чтоб нельзя было менять чужие

	admin := route.Group("admin")
	{
		admin.Post("admin/sign-up", h.AdminSignUp)
		admin.Post("admin/ban-user", mw.AdminAuth, h.BanUser)
		admin.Post("admin/create-invite-token", mw.AdminAuth, h.CreateInviteToken)
		admin.Get("admin/get-list-of-users", mw.AdminAuth, h.GetListOfUsers)
	}
}
