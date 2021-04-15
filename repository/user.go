package repository

import (
	"errors"
	"github.com/westlife0615/chak-server/config"
	"github.com/westlife0615/chak-server/model"
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

func (r *exampleUserRepository) Update(todo model.User) error {
	id := todo.Id
	_, ok := r.todos[id]
	if !ok {
		return errors.New("Has no item to update. Id - " + string(id))
	}
	r.todos[id] = todo
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

func (r *exampleUserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := config.Database.Find(&users).Error; err != nil {
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
