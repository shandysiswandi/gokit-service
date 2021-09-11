package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/pkg/uuid"
	"github.com/shandysiswandi/gokit-service/service"
)

type Endpoints struct {
	GetAllTodo  endpoint.Endpoint
	GetTodoByID endpoint.Endpoint
	CreateTodo  endpoint.Endpoint
	UpdateTodo  endpoint.Endpoint
	DeleteTodo  endpoint.Endpoint
}

func NewEndpoints(srv service.TodoService, secret string) Endpoints {
	return Endpoints{
		GetAllTodo:  JWTMiddleware(secret)(makeGetAllTodo(srv)),
		GetTodoByID: JWTMiddleware(secret)(makeGetTodoByID(srv)),
		CreateTodo:  JWTMiddleware(secret)(makeCreateTodo(srv)),
		UpdateTodo:  JWTMiddleware(secret)(makeUpdateTodo(srv)),
		DeleteTodo:  JWTMiddleware(secret)(makeDeleteTodo(srv)),
	}
}

func makeGetAllTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.GetAllTodoTodoRequest)

		data, err := s.GetAllTodo(ctx, req)
		if err != nil {
			return nil, err
		}

		return entity.GetAllTodoTodoResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: "Success get all todo",
			Data:    data,
		}, nil
	}
}

func makeGetTodoByID(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.GetTodoByIDTodoRequest)

		// validate req
		if ok := uuid.IsValidUUID(req.ID); !ok {
			return nil, errors.New("your request id is malformat")
		}

		data, err := s.GetTodoByID(ctx, req)
		if err != nil {
			return nil, err
		}

		return entity.GetTodoByIDTodoResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: "Success get todo by id",
			Data:    data,
		}, nil
	}
}

func makeCreateTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.CreateTodoRequest)

		if err := s.CreateTodo(ctx, req); err != nil {
			return nil, err
		}

		return entity.CreateTodoResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: "Success create todo",
		}, nil
	}
}

func makeUpdateTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.UpdateTodoRequest)

		if err := s.UpdateTodo(ctx, req); err != nil {
			return nil, err
		}

		return entity.UpdateTodoResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: "Success update todo",
		}, nil
	}
}

func makeDeleteTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.DeleteTodoRequest)

		if err := s.DeleteTodo(ctx, req); err != nil {
			return nil, err
		}

		return entity.DeleteTodoResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: "Success delete todo",
		}, nil
	}
}
