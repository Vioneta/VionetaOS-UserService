package v2

import codegen "github.com/Vioneta/VionetaOS-UserService/codegen/user_service"

type UserService struct{}

func NewUserService() codegen.ServerInterface {
	return &UserService{}
}
