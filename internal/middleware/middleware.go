package middleware

import (
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorGateway/pkg/consts"
	authClient "github.com/glamostoffer/ValinorProtos/auth"
	adminProto "github.com/glamostoffer/ValinorProtos/auth/admin_auth"
	clientProto "github.com/glamostoffer/ValinorProtos/auth/client_auth"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	cfg  Config
	auth *authClient.Connector
}

func New(
	//cfg Config,
	authClient *authClient.Connector,
) Middleware {
	return Middleware{
		//cfg:  cfg,
		auth: authClient,
	}
}

func (m *Middleware) ClientAuth(c *fiber.Ctx) error {
	ctx := c.Context()

	resp, err := m.auth.ClientAuth.ClientAuth(ctx, &clientProto.ClientAuthRequest{
		AccessToken: c.Cookies(consts.AccessTokenCookie),
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	c.Locals(consts.UserLocalsKey, model.UserLocals{
		UserID: resp.GetUserID(),
		Login:  resp.GetLogin(),
		Role:   resp.GetRole(),
	})

	return c.Next()
}

func (m *Middleware) AdminAuth(c *fiber.Ctx) error {
	ctx := c.Context()

	resp, err := m.auth.AdminAuth.AdminAuth(ctx, &adminProto.AdminAuthRequest{
		AccessToken: c.Cookies(consts.AccessTokenCookie),
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	c.Locals(consts.UserLocalsKey, model.UserLocals{
		UserID: resp.GetUserID(),
		Login:  resp.GetLogin(),
		Role:   resp.GetRole(),
	})

	return c.Next()
}
