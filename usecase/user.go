package usecase

import (
	"errors"
	"github.com/westlife0615/chak-server/repository"

	"github.com/westlife0615/chak-server/model"
	//"github.com/shharn/golang-clean-architecture/repository"
)

var (
	InvalidIdError = errors.New("The Id should not be 0")
)

type UserUsecase interface {
	CreateUser(model.User) error
	UpdateUser(model.User) error
	DeleteUser(int) error
	GetAllUsers() []model.User
}

type exampleUserUsecase struct {
	userRepository repository.UserRepository
}

func (s *exampleUserUsecase) CreateUser(curUser model.User) error {
	return s.userRepository.Create(curUser)
}

func (s *exampleUserUsecase) UpdateUser(curUser model.User) error {
	if curUser.Id < 1 {
		return InvalidIdError
	}
	return s.userRepository.Update(curUser)
}

func (s *exampleUserUsecase) DeleteUser(id int) error {
	if id < 1 {
		return InvalidIdError
	}
	return s.userRepository.Delete(id)
}

func (s *exampleUserUsecase) GetAllUsers() []model.User {
	todos, err := s.userRepository.GetAll()
	if err != nil {
		return []model.User{}
	}
	return todos
}


func NewUserUsecase(tr repository.UserRepository) UserUsecase {
	return &exampleUserUsecase{tr}
}