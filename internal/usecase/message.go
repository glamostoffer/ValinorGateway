package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorGateway/internal/model"
	"github.com/glamostoffer/ValinorGateway/utils/convert"
	client "github.com/glamostoffer/ValinorProtos/chat/client_chat"
)

func (uc *useCase) GetMessagesFromRoom(ctx context.Context, roomID int64) (messages []model.Message, err error) {
	messages = make([]model.Message, 0)

	out, err := uc.chat.ClientChat.GetMessagesFromRoom(ctx, &client.GetMessagesFromRoomRequest{RoomID: roomID})
	if err != nil {
		return messages, err
	}

	return convert.MessagesFromProto(out.GetMessages()), nil
}
