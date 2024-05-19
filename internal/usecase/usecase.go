package usecase

import (
	authclient "github.com/glamostoffer/ValinorProtos/auth"
)

type useCase struct {
	auth *authclient.Connector
}

func New(auth *authclient.Connector) UseCase {
	return &useCase{auth: auth}
}
