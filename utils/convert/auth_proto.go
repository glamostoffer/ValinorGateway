package convert

import (
	"github.com/glamostoffer/ValinorGateway/internal/model"
	admin "github.com/glamostoffer/ValinorProtos/auth/admin_auth"
)

func ListOfUsersFromProto(request *admin.GetListOfUsersResponse) model.GetListOfUsersResponse {
	resp := model.GetListOfUsersResponse{}
	users := request.GetUsers()

	resp.HasNext = request.GetHasNext()
	resp.Items = make([]model.User, len(users), len(users))

	for i, user := range users {
		resp.Items[i] = model.User{
			ID:        user.Id,
			Login:     user.Login,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return resp
}
