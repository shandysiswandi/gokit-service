package service

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) CreateTodo(ctx context.Context, req entity.CreateTodoRequest) error {
	return ts.dbRW.CreateTodo(ctx, entity.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.DRAFT,
	})
}

func (m *middleware) CreateTodo(ctx context.Context, req entity.CreateTodoRequest) error {
	level.Info(m.logger).Log("method", "CreateTodo", "request", req)
	return m.next.CreateTodo(ctx, req)
}
