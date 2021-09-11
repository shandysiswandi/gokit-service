package service

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) GetAllTodo(ctx context.Context, req entity.GetAllTodoTodoRequest) ([]entity.Todo, error) {
	data, err := ts.dbRW.GetAllTodo(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *middleware) GetAllTodo(ctx context.Context, req entity.GetAllTodoTodoRequest) ([]entity.Todo, error) {
	level.Info(m.logger).Log("method", "GetAllTodo", "request", req)
	return m.next.GetAllTodo(ctx, req)
}
