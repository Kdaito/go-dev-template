package response

import "go-dev-sample/internal/domain/model"

type GetUserByIdResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGetUserByIdResponse(userModel *model.User) *GetUserByIdResponse {
	return &GetUserByIdResponse{
		Id:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
}
