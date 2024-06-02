package model

type (
	Message struct {
		RoomID   int64  `json:"roomID"`
		ClientID int64  `json:"clientID"`
		Content  string `json:"content"`
		SentAt   int64  `json:"sentAt"` // unix
		Username string `json:"username"`
	}

	GetMessagesFromRoomRequest struct {
		RoomID int64 `json:"roomID"`
	}

	GetMessagesFromRoomResponse struct {
		Messages []Message `json:"items"`
	}
)
