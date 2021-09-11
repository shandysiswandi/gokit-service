package service

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (ts *todoService) GetAllTodo(ctx context.Context, req entity.GetAllTodoTodoRequest) ([]entity.Todo, error) {
	return []entity.Todo{}, nil
}
