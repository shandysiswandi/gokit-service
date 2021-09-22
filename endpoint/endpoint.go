package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/pkg/uuid"
	"github.com/shandysiswandi/gokit-service/service"
)

type Endpoints interface {
	GetAllTodo(ctx context.Context, request interface{}) (response interface{}, err error)
	GetTodoByID(ctx context.Context, request interface{}) (response interface{}, err error)
	CreateTodo(ctx context.Context, request interface{}) (response interface{}, err error)
	UpdateTodo(ctx context.Context, request interface{}) (response interface{}, err error)
	DeleteTodo(ctx context.Context, request interface{}) (response interface{}, err error)
}

type endpoints struct {
	jwtSecret string
	service   service.TodoService
}

func NewEndpoints(srv service.TodoService, secret string) Endpoints {
	return &endpoints{
		jwtSecret: secret,
		service:   srv,
	}
}

func (e *endpoints) GetAllTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(entity.GetAllTodoRequest)

	data, err := e.service.GetAllTodo(ctx, req)
	if err != nil {
		return nil, err
	}

	return entity.GetAllTodoResponse{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success get all todo",
		Data:    data,
	}, nil
}

func (e *endpoints) GetTodoByID(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(entity.GetTodoByIDRequest)

	// validate req
	if ok := uuid.IsValidUUID(req.ID); !ok {
		return nil, errors.New("your request id is malformat")
	}

	data, err := e.service.GetTodoByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return entity.GetTodoByIDResponse{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success get todo by id",
		Data:    data,
	}, nil
}

func (e *endpoints) CreateTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(entity.CreateTodoRequest)

	if err := e.service.CreateTodo(ctx, req); err != nil {
		return nil, err
	}

	return entity.CreateTodoResponse{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success create todo",
	}, nil
}

func (e *endpoints) UpdateTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(entity.UpdateTodoRequest)

	if err := e.service.UpdateTodo(ctx, req); err != nil {
		return nil, err
	}

	return entity.UpdateTodoResponse{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success update todo",
	}, nil
}

func (e *endpoints) DeleteTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(entity.DeleteTodoRequest)

	if err := e.service.DeleteTodo(ctx, req); err != nil {
		return nil, err
	}

	return entity.DeleteTodoResponse{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success delete todo",
	}, nil
}
