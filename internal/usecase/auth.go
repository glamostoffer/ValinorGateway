package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorGateway/utils/convert"
	admin "github.com/glamostoffer/ValinorProtos/auth/admin_auth"
	client "github.com/glamostoffer/ValinorProtos/auth/client_auth"
)

// ==================== ADMIN ==================== //

func (uc *useCase) AdminAuth(
	ctx context.Context,
	req model.AdminAuthRequest,
) (resp model.AdminAuthResponse, err error) {
	out, err := uc.auth.AdminAuth.AdminAuth(
		ctx,
		&admin.AdminAuthRequest{
			AccessToken: req.AccessToken,
		},
	)
	if err != nil {
		return resp, err
	}

	return model.AdminAuthResponse{
		UserID: out.GetUserID(),
		Login:  out.GetLogin(),
		Role:   out.GetRole(),
	}, nil
}

func (uc *useCase) AdminSignUp(ctx context.Context, req model.AdminSignUpRequest) error {
	_, err := uc.auth.AdminAuth.AdminSignUp(
		ctx,
		&admin.AdminSignUpRequest{
			Login:       req.Login,
			Password:    req.Password,
			InviteToken: req.InviteToken,
		},
	)

	return err
}

func (uc *useCase) BanUser(ctx context.Context, req model.BanUserRequest) error {
	_, err := uc.auth.AdminAuth.BanUser(
		ctx,
		&admin.BanUserRequest{
			UserID: req.UserID,
		},
	)

	return err
}

func (uc *useCase) GetListOfUsers(
	ctx context.Context,
	req model.GetListOfUsersRequest,
) (resp model.GetListOfUsersResponse, err error) {
	out, err := uc.auth.AdminAuth.GetListOfUsers(
		ctx,
		&admin.GetListOfUsersRequest{
			Limit:  req.Limit,
			Offset: req.Offset,
		},
	)
	if err != nil {
		return resp, err
	}

	return convert.ListOfUsersFromProto(out), nil
}

func (uc *useCase) CreateInviteToken(
	ctx context.Context,
	req model.CreateInviteTokenRequest,
) (resp model.CreateInviteTokenResponse, err error) {
	out, err := uc.auth.AdminAuth.CreateInviteToken(
		ctx,
		&admin.CreateInviteTokenRequest{
			Ttl: req.TTL,
		},
	)
	if err != nil {
		return resp, err
	}

	return model.CreateInviteTokenResponse{
		Token: out.GetToken(),
	}, nil
}

// ==================== USER ==================== //

func (uc *useCase) ClientAuth(
	ctx context.Context,
	req model.ClientAuthRequest,
) (resp model.ClientAuthResponse, err error) {
	out, err := uc.auth.ClientAuth.ClientAuth(
		ctx,
		&client.ClientAuthRequest{
			AccessToken: req.AccessToken,
		},
	)
	if err != nil {
		return resp, err
	}

	return model.ClientAuthResponse{
		UserID: out.GetUserID(),
		Login:  out.GetLogin(),
		Role:   out.GetRole(),
	}, nil
}

func (uc *useCase) SignUp(ctx context.Context, req model.SignUpRequest) error {
	_, err := uc.auth.ClientAuth.SignUp(
		ctx,
		&client.SignUpRequest{
			Login:    req.Login,
			Password: req.Password,
		},
	)

	return err
}

func (uc *useCase) SignIn(
	ctx context.Context,
	req model.SignInRequest,
) (resp model.SignInResponse, err error) {
	out, err := uc.auth.ClientAuth.SignIn(
		ctx,
		&client.SignInRequest{
			Login:    req.Login,
			Password: req.Password,
		},
	)
	if err != nil {
		return resp, err
	}

	return model.SignInResponse{
		Token: out.GetToken(),
	}, nil
}

func (uc *useCase) UpdateClientDetails(ctx context.Context, req model.UpdateClientDetailsRequest) error {
	_, err := uc.auth.ClientAuth.UpdateClientDetails(
		ctx,
		&client.UpdateClientDetailsRequest{
			ClientID: req.ClientID,
			Username: req.Login,
			Password: req.Password,
		},
	)

	return err
}

func (uc *useCase) GetClientDetails(
	ctx context.Context,
	req model.GetClientDetailsRequest,
) (resp model.GetClientDetailsResponse, err error) {
	out, err := uc.auth.ClientAuth.GetClientDetails(
		ctx,
		&client.GetClientDetailsRequest{
			ClientID: req.ClientID,
		},
	)
	if err != nil {
		return resp, err
	}

	return model.GetClientDetailsResponse{
		Login:     out.GetUsername(),
		Role:      out.GetRole(),
		CreatedAt: out.GetCreatedAt(),
		UpdatedAt: out.GetUpdatedAt(),
	}, nil
}
