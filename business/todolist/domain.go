package todolist

import "time"

type Todo struct {
	Id        int
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
