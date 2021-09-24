package service

import (
	"context"
	"fmt"

	"github.com/go-kit/log/level"
	"github.com/shandysiswandi/gokit-service/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ts *todoService) GetTodoByID(ctx context.Context, req entity.GetTodoByIDRequest) (entity.Todo, error) {
	data, err := ts.dbRW.GetTodoByID(ctx, req.ID)
	if err != nil {
		return entity.Todo{}, status.Error(codes.NotFound, fmt.Sprintf("todo with id %s is not found", req.ID))
	}
	return data, nil
}

func (m *middleware) GetTodoByID(ctx context.Context, req entity.GetTodoByIDRequest) (entity.Todo, error) {
	level.Info(m.logger).Log("method", "GetTodoByID", "request", req)
	return m.next.GetTodoByID(ctx, req)
}
