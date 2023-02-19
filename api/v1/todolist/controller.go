package todolist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mashbens/todolist/api/common/obj"
	_response "github.com/mashbens/todolist/api/common/response"
	"github.com/mashbens/todolist/api/v1/todolist/request"
	"github.com/mashbens/todolist/api/v1/todolist/resp"
	service "github.com/mashbens/todolist/business/todolist"
)

type TodoListController struct {
	TodoListService service.TodoListService
}

func NewTodoListController(
	TodoListService service.TodoListService,

) *TodoListController {
	return &TodoListController{
		TodoListService: TodoListService,
	}
}

var validate *validator.Validate

func (controller *TodoListController) CreateTodo(c echo.Context) error {
	payload := new(request.TodoRequest)
	err := c.Bind(payload)

	validate = validator.New()
	err = validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
	if payload.Title == "" {
		response := _response.BuildErrorResponse("Bad Request", "title cannot be null")
		return c.JSON(http.StatusBadRequest, response)
	}
	if payload.Email == "" || err != nil {
		response := _response.BuildErrorResponse("Bad Request", "email cannot be null")
		return c.JSON(http.StatusBadRequest, response)
	}
	res, err := controller.TodoListService.InsertTodoList(request.ToService(*payload))
	if err != nil {
		fmt.Println(err)
	}
	data := resp.FromService(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *TodoListController) FindAll(c echo.Context) error {
	res, _ := controller.TodoListService.FindAll()
	data := resp.FromeServiceSlice(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *TodoListController) FindTodoById(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	res, err := controller.TodoListService.FindTodoById(intID)
	if err != nil {
		response := _response.BuildErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *TodoListController) UpdateTodoById(c echo.Context) error {
	payload := new(request.TodoRequest)
	err := c.Bind(payload)

	if payload.Title == "" {
		response := _response.BuildErrorResponse("Bad Request", "title cannot be null")
		return c.JSON(http.StatusBadRequest, response)
	}

	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	res, err := controller.TodoListService.UpdateTodoById(intID, request.ToService(*payload))
	if err != nil {
		response := _response.BuildErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *TodoListController) DeleteTodoById(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	err := controller.TodoListService.DeleteTodoById(intID)
	if err != nil {
		response := _response.BuildErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
		return c.JSON(http.StatusBadRequest, response)
	}
	response := _response.BuildSuccsessResponse("Succses", obj.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
