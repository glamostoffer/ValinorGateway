package model

type UserLocals struct {
	UserID int64  `json:"userID"`
	Login  string `json:"login"`
	Role   string `json:"role"`
}

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"createdAt"` // time.Time.Unix()
	UpdatedAt int64  `json:"updatedAt"`
}

// ==================== ADMIN ==================== //

type (
	AdminAuthRequest struct {
		AccessToken string `json:"accessToken" validate:"required"`
	}
	AdminAuthResponse struct {
		UserID int64  `json:"userID"`
		Login  string `json:"login"`
		Role   string `json:"role"`
	}

	AdminSignUpRequest struct {
		Login       string `json:"login" validate:"required"`
		Password    string `json:"password" validate:"required"`
		InviteToken string `json:"inviteToken" validate:"required"`
	}

	BanUserRequest struct {
		UserID int64 `json:"userID" validate:"required"`
	}

	CreateInviteTokenRequest struct {
		TTL int64 `json:"TTL" validate:"required"`
	}
	CreateInviteTokenResponse struct {
		Token string `json:"token"`
	}

	GetListOfUsersRequest struct {
		Limit  int64 `json:"limit" validate:"min=0"`
		Offset int64 `json:"offset" validate:"max=228,min=0"`
	}
	GetListOfUsersResponse struct {
		Items   []User `json:"items"`
		HasNext bool   `json:"hasNext"`
	}
)

// ==================== USER ==================== //

type (
	ClientAuthRequest struct {
		AccessToken string `json:"accessToken" validate:"required"`
	}
	ClientAuthResponse struct {
		UserID int64  `json:"userID"`
		Login  string `json:"login"`
		Role   string `json:"role"`
	}

	SignUpRequest struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	SignInRequest struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	SignInResponse struct {
		Token string `json:"token"`
	}

	GetClientDetailsRequest struct {
		ClientID int64 `json:"clientID" validate:"required"`
	}
	GetClientDetailsResponse struct {
		Login     string `json:"login"`
		Role      string `json:"role"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
	}

	UpdateClientDetailsRequest struct {
		ClientID int64   `json:"clientID" validate:"required"`
		Login    *string `json:"login"`
		Password *string `json:"password"`
	}
)
