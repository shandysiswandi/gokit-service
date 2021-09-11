package service

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) GetTodoByID(ctx context.Context, req entity.GetTodoByIDTodoRequest) (entity.Todo, error) {
	data, err := ts.dbRW.GetTodoByID(ctx, req.ID)
	if err != nil {
		return entity.Todo{}, err
	}
	return data, nil
}

func (m *middleware) GetTodoByID(ctx context.Context, req entity.GetTodoByIDTodoRequest) (entity.Todo, error) {
	level.Info(m.logger).Log("method", "GetTodoByID", "request", req)
	return m.next.GetTodoByID(ctx, req)
}
