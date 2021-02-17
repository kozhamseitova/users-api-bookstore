package services

import (
	users "gitlab.com/tleuzhan13/bookstore/users-api/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"gitlab.com/tleuzhan13/bookstore/users-api/utils/crypto"
	"gitlab.com/tleuzhan13/bookstore/users-api/utils/date"
)

type UsersService struct{}

type UsersServiceInterface interface {
	GetUser(int64) (*users.User, rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *UsersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	dao := &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *UsersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBFormat()
	user.Password = crypto.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *UsersService) DeleteUser(userId int64) rest_errors.RestErr {
	dao := &users.User{Id: userId}
	return dao.Delete()
}

func (s *UsersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *UsersService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
