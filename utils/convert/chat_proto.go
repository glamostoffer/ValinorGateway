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

func MessagesFromProto(protoMessages []*client_chat.Message) []model.Message {
	messages := make([]model.Message, 0, len(protoMessages))

	for _, msg := range protoMessages {
		messages = append(messages, model.Message{
			RoomID:   msg.RoomID,
			ClientID: msg.ClientID,
			Content:  msg.Message,
			SentAt:   msg.SentAt,
			Username: msg.Username,
		})
	}

	return messages
}
