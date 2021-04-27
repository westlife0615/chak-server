package repository

import (
	"errors"
	"github.com/westlife0615/chak-server/config"
	"github.com/westlife0615/chak-server/model"
	"time"
)

type UserRepository interface {
	Create(model.User) error
	Update(model.User) error
	Delete(int) error
	GetAll() ([]model.User, error)
}

type exampleUserRepository struct {
	todos  map[int]model.User
	nextId int
}

func (r *exampleUserRepository) Create(curUser model.User) error {
	var user model.User = curUser
	config.Database.Create(&user)
	return nil

}

func (r *exampleUserRepository) Update(curUser model.User) error {
	//1  find one
	var foundUser model.User

	if foundUserError := config.Database.First(&foundUser, curUser.Id).Error; foundUserError != nil {
		return foundUserError
	}

	// 2 update model
	foundUser.Email = curUser.Email
	foundUser.UpdatedAt = time.Now()

	// 3 save
	if saveUserError := config.Database.Save(&foundUser).Error; saveUserError != nil {
		return saveUserError
	}
	return nil
}

func (r *exampleUserRepository) Delete(id int) error {
	_, ok := r.todos[id]
	if !ok {
		return errors.New("Has no item to delete. Id - " + string(id))
	}
	delete(r.todos, id)
	return nil
}

// TODO : omit Password
func (r *exampleUserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	//if err := config.Database.Find(&users).Error; err != nil {
	//	return users, err
	//}
	if err := config.Database.Table("users").Select([]string{"id", "email"}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

/*
"" string
'' rural
*/
func NewUserRepository() UserRepository {
	return &exampleUserRepository{
		//todos: map[int]model.User{
		//	1: model.User{1, "123123", "bigin@bigin.io", time.Now(), time.Now()},
		//	2: model.User{2, "123123", "westlife0615@bigin.io", time.Now(), time.Now()},
		//},
		//nextId: 3,
	}
}
