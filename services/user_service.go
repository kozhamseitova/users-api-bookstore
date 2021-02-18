package services

import (
	users "gitlab.com/tleuzhan13/bookstore/users-api/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"gitlab.com/tleuzhan13/bookstore/users-api/utils/crypto"
	"gitlab.com/tleuzhan13/bookstore/users-api/utils/date"
)

type UsersService struct {
	Repository Repository
}

type Repository interface {
	Get(id int64) (*users.User, rest_errors.RestErr)
	Save(user *users.User) (*users.User, rest_errors.RestErr)
	Update(user *users.User) (*users.User, rest_errors.RestErr)
	Delete(id int64) rest_errors.RestErr
	FindByStatus(status string) ([]*users.User, rest_errors.RestErr)
	FindByEmailAndPassword(email string, password string) (*users.User, rest_errors.RestErr)
}

func (s *UsersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	user, err := s.Repository.Get(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UsersService) CreateUser(user *users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBFormat()
	user.Password = crypto.GetMd5(user.Password)
	userSaved, err := s.Repository.Save(user)
	if err != nil {
		return nil, err
	}
	return userSaved, nil
}

func (s *UsersService) UpdateUser(isPartial bool, user *users.User) (*users.User, rest_errors.RestErr) {

	current, err := s.Repository.Get(user.Id)
	if err != nil {
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
	current, err = s.Repository.Update(current)
	if err != nil {
		return nil, err
	}
	return current, nil
}

func (s *UsersService) DeleteUser(userId int64) rest_errors.RestErr {
	return s.Repository.Delete(userId)
}

func (s *UsersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	return s.Repository.FindByStatus(status)
}

func (s *UsersService) LoginUser(request *users.LoginRequest) (*users.User, rest_errors.RestErr) {
	usr, err := s.Repository.FindByEmailAndPassword(request.Email, crypto.GetMd5(request.Password))
	if err != nil {
		return nil, err
	}
	return usr, nil
}
