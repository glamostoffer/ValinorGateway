package convert

import (
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorProtos/chat/client_chat"
)

func RoomsListFromProto(protoRooms []*client_chat.Room) []model.Room {
	rooms := make([]model.Room, 0, len(protoRooms))

	for _, room := range protoRooms {
		rooms = append(rooms, model.Room{
			ID:        room.GetRoomID(),
			Name:      room.GetName(),
			OwnerID:   room.GetOwnerID(),
			ClientIDs: room.GetClientIDs(),
		})
	}

	return rooms
}
