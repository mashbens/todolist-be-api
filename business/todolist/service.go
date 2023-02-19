package todolist

import (
	"errors"
)

type TodoListRepo interface {
	FindAll() ([]Todo, error)
	InsertTodoList(Todo) (Todo, error)
	FindTodoById(id int) (Todo, error)
	UpdateTodoById(int, Todo) (Todo, error)
	DeleteTodoById(id int) error
}

type TodoListService interface {
	FindAll() ([]Todo, error)
	InsertTodoList(Todo) (Todo, error)
	FindTodoById(id int) (Todo, error)
	UpdateTodoById(int, Todo) (Todo, error)
	DeleteTodoById(id int) error
}

type todoListService struct {
	todoListRepo TodoListRepo
}

func NewTodolistService(
	todoListRepo TodoListRepo,
) TodoListService {
	return &todoListService{
		todoListRepo: todoListRepo}
}

func (c *todoListService) FindAll() (res []Todo, err error) {
	t, err := c.todoListRepo.FindAll()
	if err != nil {
		return res, err
	}
	return t, nil
}

func (c *todoListService) InsertTodoList(todo Todo) (res Todo, err error) {
	t, err := c.todoListRepo.InsertTodoList(todo)
	if err != nil {
		return res, err
	}
	return t, nil
}

func (c *todoListService) FindTodoById(id int) (res Todo, err error) {
	t, err := c.todoListRepo.FindTodoById(id)
	if err != nil {
		return res, err
	}
	return t, nil
}
func (c *todoListService) UpdateTodoById(id int, todo Todo) (res Todo, err error) {
	t, _ := c.todoListRepo.FindTodoById(id)
	if t.Id == 0 {
		return res, errors.New("id not found")
	}
	update, err := c.todoListRepo.UpdateTodoById(id, todo)
	if err != nil {
		return res, err
	}
	update.Id = id
	return update, nil
}

func (c *todoListService) DeleteTodoById(id int) error {
	t, _ := c.todoListRepo.FindTodoById(id)
	if t.Id == 0 {
		return errors.New("id not found")
	}
	err := c.todoListRepo.DeleteTodoById(t.Id)
	if err != nil {
		return errors.New("error deleting")
	}
	return nil
}
