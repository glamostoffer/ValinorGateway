package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorGateway/utils/convert"
	auth "github.com/glamostoffer/ValinorProtos/auth/admin_auth"
	client "github.com/glamostoffer/ValinorProtos/chat/client_chat"
)

func (uc *useCase) CreateRoom(ctx context.Context, req model.CreateRoomRequest) (roomID int64, err error) {
	out, err := uc.chat.ClientChat.CreateRoom(ctx, &client.CreateRoomRequest{
		Name:     req.Name,
		ClientID: req.ClientID,
	})
	if err != nil {
		return -1, err
	}

	return out.GetRoomID(), err
}

func (uc *useCase) GetListOfRooms(ctx context.Context, clientID int64) (rooms []model.Room, err error) {
	rooms = make([]model.Room, 0)

	out, err := uc.chat.ClientChat.GetListOfRooms(ctx, &client.GetListOfRoomsRequest{ClientID: clientID})
	if err != nil {
		return rooms, err
	}

	return convert.RoomsListFromProto(out.GetRooms()), nil
}

func (uc *useCase) AddClientToRoom(ctx context.Context, req model.AddClientToRoomRequest) (err error) {
	authResponse, err := uc.auth.AdminAuth.GetClientIDByLogin(ctx, &auth.GetClientIDByLoginRequest{Login: req.ClientLogin})
	if err != nil {
		return err
	}

	_, err = uc.chat.ClientChat.AddClientToRoom(ctx, &client.AddClientToRoomRequest{
		RoomID:   req.RoomID,
		ClientID: authResponse.GetClientID(),
	})

	return err
}

func (uc *useCase) RemoveClientFromRoom(ctx context.Context, req model.RemoveClientFromRoomRequest) (err error) {
	_, err = uc.chat.ClientChat.RemoveClientFromRoom(ctx, &client.RemoveClientFromRoomRequest{
		RoomID:   req.RoomID,
		ClientID: req.ClientID,
	})

	return err
}
