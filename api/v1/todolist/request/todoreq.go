package request

import "github.com/mashbens/todolist/business/todolist"

type TodoRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func ToService(t TodoRequest) todolist.Todo {
	return todolist.Todo{
		Title: t.Title,
		Email: t.Email,
	}
}
