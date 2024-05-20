package usecase

import (
	authclient "github.com/glamostoffer/ValinorProtos/auth"
	chatclient "github.com/glamostoffer/ValinorProtos/chat"
)

type useCase struct {
	auth *authclient.Connector
	chat *chatclient.Connector
}

func New(auth *authclient.Connector, chat *chatclient.Connector) UseCase {
	return &useCase{
		auth: auth,
		chat: chat,
	}
}
