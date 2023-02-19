package resp

import (
	"time"

	"github.com/mashbens/todolist/business/todolist"
)

type TodoResp struct {
	Id        int
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromService(t todolist.Todo) TodoResp {
	return TodoResp{
		Id:        t.Id,
		Title:     t.Title,
		Email:     t.Email,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func FromeServiceSlice(data []todolist.Todo) []TodoResp {
	var todoListArray []TodoResp
	for key := range data {
		todoListArray = append(todoListArray, FromService(data[key]))
	}
	return todoListArray
}
