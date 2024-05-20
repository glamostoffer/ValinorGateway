package delivery

import (
	"github.com/glamostoffer/ValinorGateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(route fiber.Router, mw middleware.Middleware, h *Handler) {
	route.Post("sign-up", h.SignUp)
	route.Post("sign-in", h.SignIn)

	client := route.Group("client")
	{
		client.Post("details/get", mw.ClientAuth, h.GetClientDetails)
		client.Post("details/update", mw.ClientAuth, h.UpdateClientDetails)
	}

	admin := route.Group("admin")
	{
		admin.Post("sign-up", h.AdminSignUp)
		admin.Post("ban-user", mw.AdminAuth, h.BanUser)
		admin.Post("create-invite-token", mw.AdminAuth, h.CreateInviteToken)
		admin.Post("get-list-of-users", mw.AdminAuth, h.GetListOfUsers)
	}

	chat := route.Group("chat")
	{
		room := chat.Group("room")
		{
			room.Post("create", mw.ClientAuth, h.CreateRoom)
			room.Post("get", mw.ClientAuth, h.GetClientRooms)
			room.Post("add", mw.ClientAuth, h.AddClientToRoom)
			room.Post("remove", mw.ClientAuth, h.RemoveClientFromRoom)
		}
	}
}
