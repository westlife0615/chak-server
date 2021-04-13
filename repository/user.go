package repository

import (
	"errors"
	"time"
	"github.com/westlife0615/chak-server/model"

)

type UserRepository interface {
	Create(model.User) error
	Update(model.User) error
	Delete(int) error
	GetAll() ([]model.User, error)
}

type exampleUserRepository struct {
	todos map[int]model.User
	nextId int
}

func (r *exampleUserRepository) Create(todo model.User) error {
	nextId := r.nextId
	todo.Id = nextId
	r.todos[r.nextId] = todo
	r.nextId = nextId + 1
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
	arr := []model.User{}
	for _, todo := range r.todos {
		arr = append(arr, todo)
	}
	return arr, nil
}

func NewUserRepository() UserRepository {
	return &exampleUserRepository{
		todos: map[int]model.User{
			1: model.User{1, "Content 1", time.Now(), time.Now()},
			2: model.User{2, "Content 2", time.Now(), time.Now()},
		},
		nextId: 3,
	}
}