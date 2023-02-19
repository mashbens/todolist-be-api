package todolist

import (
	"time"

	"github.com/mashbens/todolist/business/todolist"
)

type Todo struct {
	Id        int
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Todo) toService() todolist.Todo {
	return todolist.Todo{
		Id:        t.Id,
		Title:     t.Title,
		Email:     t.Email,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func fromService(t todolist.Todo) Todo {
	return Todo{
		Id:        t.Id,
		Title:     t.Title,
		Email:     t.Email,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
func toServiceList(data []Todo) []todolist.Todo {
	a := []todolist.Todo{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
