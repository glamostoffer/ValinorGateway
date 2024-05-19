package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorGateway/internal/model"
)

type UseCase interface {
	Auth
}

type Auth interface {

	// ==================== ADMIN ==================== //

	AdminAuth(ctx context.Context, req model.AdminAuthRequest) (model.AdminAuthResponse, error)
	AdminSignUp(ctx context.Context, req model.AdminSignUpRequest) error
	BanUser(ctx context.Context, req model.BanUserRequest) error
	CreateInviteToken(ctx context.Context, req model.CreateInviteTokenRequest) (model.CreateInviteTokenResponse, error)
	GetListOfUsers(ctx context.Context, req model.GetListOfUsersRequest) (model.GetListOfUsersResponse, error)

	// ==================== USER ==================== //

	ClientAuth(ctx context.Context, req model.ClientAuthRequest) (model.ClientAuthResponse, error)
	SignUp(ctx context.Context, req model.SignUpRequest) error
	SignIn(ctx context.Context, req model.SignInRequest) (model.SignInResponse, error)
	GetClientDetails(ctx context.Context, req model.GetClientDetailsRequest) (model.GetClientDetailsResponse, error)
	UpdateClientDetails(ctx context.Context, req model.UpdateClientDetailsRequest) error
}
