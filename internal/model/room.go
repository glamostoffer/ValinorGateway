package model

type (
	Room struct {
		ID        int64   `json:"roomID"`
		Name      string  `json:"name"`
		OwnerID   int64   `json:"ownerID"`
		ClientIDs []int64 `json:"clientIDs"`
	}

	CreateRoomRequest struct {
		Name     string `json:"name" validate:"required"`
		ClientID int64  `json:"-"`
	}

	GetListOfRoomsResponse struct {
		Rooms []Room `json:"items"`
	}

	AddClientToRoomRequest struct {
		RoomID   int64 `json:"roomID" validate:"required"`
		OwnerID  int64 `json:"ownerID" validate:"required"`
		ClientID int64 `json:"clientID" validate:"required"`
	}

	RemoveClientFromRoomRequest struct {
		RoomID   int64 `json:"roomID" validate:"required"`
		OwnerID  int64 `json:"ownerID" validate:"required"`
		ClientID int64 `json:"clientID" validate:"required"`
	}
)
