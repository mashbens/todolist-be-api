package todolist

import (
	"github.com/mashbens/todolist/business/todolist"
	"gorm.io/gorm"
)

type TodoListRepo struct {
	db *gorm.DB
}

func NewTodoListRepo(db *gorm.DB) todolist.TodoListRepo {
	return &TodoListRepo{
		db: db,
	}
}

func (r *TodoListRepo) FindAll() (data []todolist.Todo, err error) {
	record := []Todo{}
	res := r.db.Find(&record).Debug()
	if res.Error != nil {
		return nil, res.Error
	}
	return toServiceList(record), nil
}

func (r *TodoListRepo) InsertTodoList(data todolist.Todo) (todolist.Todo, error) {
	record := fromService(data)
	res := r.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *TodoListRepo) FindTodoById(id int) (todolist.Todo, error) {
	var record Todo
	res := r.db.Where("id = ?", id).First(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *TodoListRepo) UpdateTodoById(id int, data todolist.Todo) (todolist.Todo, error) {
	record := fromService(data)
	res := r.db.Where("id = ?", id).Updates(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	res = r.db.Where("id = ?", id).Find(&record)
	return record.toService(), nil
}
func (r *TodoListRepo) DeleteTodoById(id int) error {
	var record Todo
	res := r.db.Where("id = ?", id).Delete(&record)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
