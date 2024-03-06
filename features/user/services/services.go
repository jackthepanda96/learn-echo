package services

import "21-api/features/user"

type service struct {
	model user.UserModel
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
	}
}

func (s *service) Register(newData user.User) error {
	return nil
}
func (s *service) Login(loginData user.User) (user.User, error) {
	return user.User{}, nil
}
