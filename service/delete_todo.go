package service

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) DeleteTodo(ctx context.Context, req entity.DeleteTodoRequest) error {
	return ts.dbRW.DeleteTodo(ctx, req.ID)
}

func (m *middleware) DeleteTodo(ctx context.Context, req entity.DeleteTodoRequest) error {
	level.Info(m.logger).Log("method", "DeleteTodo", "request", req)
	return m.next.DeleteTodo(ctx, req)
}
