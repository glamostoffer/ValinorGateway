package delivery

import (
	"github.com/glamostoffer/ValinorGateway/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	vl  *validator.Validate
	cfg Config
	uc  usecase.UseCase
}

func New(
	cfg Config,
	uc usecase.UseCase,
) *Handler {
	return &Handler{
		vl:  validator.New(),
		cfg: cfg,
		uc:  uc,
	}
}

func (h *Handler) validateRequest(c *fiber.Ctx, request any) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return h.vl.StructCtx(c.Context(), request)
}
