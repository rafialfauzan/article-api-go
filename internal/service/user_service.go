package service

import (
	"errors"
	"golang-api/internal/model"
	"golang-api/internal/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
	UpdateUser(id uint, user *model.User) (model.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *model.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(id uint, user *model.User) (model.User, error) {
	if user.Name == "" {
		return model.User{}, errors.New("name is required")
	}
	if user.Email == "" {
		return model.User{}, errors.New("email is required")
	}

	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{}, errors.New("user not found")
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	if err := s.repo.Update(&existingUser); err != nil {
		return model.User{}, err
	}

	return existingUser, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}
	return s.repo.Delete(id)
}
