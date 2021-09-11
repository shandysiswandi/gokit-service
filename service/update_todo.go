package service

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) UpdateTodo(ctx context.Context, req entity.UpdateTodoRequest) error {
	return ts.dbRW.UpdateTodo(ctx, entity.Todo{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      "DONE",
	})
}

func (m *middleware) UpdateTodo(ctx context.Context, req entity.UpdateTodoRequest) error {
	level.Info(m.logger).Log("method", "UpdateTodo", "request", req)
	return m.next.UpdateTodo(ctx, req)
}
