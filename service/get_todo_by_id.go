package service

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) GetTodoByID(ctx context.Context, req entity.GetTodoByIDTodoRequest) (entity.Todo, error) {
	return entity.Todo{}, nil
}
