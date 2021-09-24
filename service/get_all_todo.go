package service

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ts *todoService) GetAllTodo(ctx context.Context, req entity.GetAllTodoRequest) ([]entity.Todo, error) {
	data, err := ts.dbRW.GetAllTodo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return data, nil
}

func (m *middleware) GetAllTodo(ctx context.Context, req entity.GetAllTodoRequest) ([]entity.Todo, error) {
	level.Info(m.logger).Log("method", "GetAllTodo", "request", req)
	return m.next.GetAllTodo(ctx, req)
}
