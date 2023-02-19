package todolist

import (
	"github.com/mashbens/todolist/business/todolist"
	"github.com/mashbens/todolist/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) todolist.TodoListRepo {
	var todoListRepository todolist.TodoListRepo

	if dbCon.Driver == util.MYSQL {
		todoListRepository = NewTodoListRepo(dbCon.MYSQL)
		dbCon.MYSQL.AutoMigrate(&Todo{})
	} else {
		panic("Database driver not supported")
	}
	return todoListRepository
}
