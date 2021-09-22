package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/shandysiswandi/gokit-service/endpoint"
	v1 "github.com/shandysiswandi/gokit-service/proto"
)

type server struct {
	getAllTodo  grpctransport.Handler
	getTodoByID grpctransport.Handler
	createTodo  grpctransport.Handler
	updateTodo  grpctransport.Handler
	deleteTodo  grpctransport.Handler
}

func NewServer(end endpoint.Endpoints) v1.TodoServiceServer {
	// options like middleware etc.
	options := []grpctransport.ServerOption{
		// grpctransport.ServerBefore(jwt.GRPCToContext()),
		// grpctransport.ServerAfter()
		// grpctransport.ServerErrorHandler()
	}

	return &server{
		getAllTodo: grpctransport.NewServer(
			end.GetAllTodo,
			decodeGetAllTodo,
			encodeGetAllTodo,
			options...,
		),
		getTodoByID: grpctransport.NewServer(
			end.GetTodoByID,
			decodeGetTodoByID,
			encodeGetTodoByID,
			options...,
		),
		createTodo: grpctransport.NewServer(
			end.CreateTodo,
			decodeCreateTodo,
			encodeCreateTodo,
			options...,
		),
		updateTodo: grpctransport.NewServer(
			end.UpdateTodo,
			decodeUpdateTodo,
			encodeUpdateTodo,
			options...,
		),
		deleteTodo: grpctransport.NewServer(
			end.DeleteTodo,
			decodeDeleteTodo,
			encodeDeleteTodo,
			options...,
		),
	}
}
