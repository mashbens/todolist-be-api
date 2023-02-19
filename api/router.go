package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mashbens/todolist/api/v1/todolist"
)

type Controller struct {
	Todolist *todolist.TodoListController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	todoRoutes := e.Group("activity-groups")
	todoRoutes.GET("", controller.Todolist.FindAll)
	todoRoutes.GET("/:id", controller.Todolist.FindTodoById)
	todoRoutes.POST("", controller.Todolist.CreateTodo)
	todoRoutes.PUT("/:id", controller.Todolist.UpdateTodoById)
	todoRoutes.DELETE("/:id", controller.Todolist.DeleteTodoById)

}
