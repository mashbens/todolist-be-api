package modules

import (
	"github.com/mashbens/todolist/api"
	todolist "github.com/mashbens/todolist/api/v1/todolist"
	todolistservice "github.com/mashbens/todolist/business/todolist"
	"github.com/mashbens/todolist/config"
	todolistrepo "github.com/mashbens/todolist/repository/todolist"
	"github.com/mashbens/todolist/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	todoListRepo := todolistrepo.RepositoryFactory(dbCon)

	todoListService := todolistservice.NewTodolistService(todoListRepo)
	controller := api.Controller{
		Todolist: todolist.NewTodoListController(todoListService),
	}
	return controller
}
