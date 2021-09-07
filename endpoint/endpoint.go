package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/service"
)

type Endpoints struct {
	GetAllTodo  endpoint.Endpoint
	GetTodoByID endpoint.Endpoint
	CreateTodo  endpoint.Endpoint
	UpdateTodo  endpoint.Endpoint
	DeleteTodo  endpoint.Endpoint
}

func NewEndpoints(srv service.TodoService) Endpoints {
	return Endpoints{
		GetAllTodo:  makeGetAllTodo(srv),
		GetTodoByID: makeGetTodoByID(srv),
		CreateTodo:  makeCreateTodo(srv),
		UpdateTodo:  makeUpdateTodo(srv),
		DeleteTodo:  makeDeleteTodo(srv),
	}
}
