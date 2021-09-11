package service

import (
	"context"
	"errors"

	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/repository"
)

type TodoService interface {
	GetAllTodo(ctx context.Context, req entity.GetAllTodoTodoRequest) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, req entity.GetTodoByIDTodoRequest) (entity.Todo, error)

	CreateTodo(ctx context.Context, req entity.CreateTodoRequest) error
	UpdateTodo(ctx context.Context, req entity.UpdateTodoRequest) error
	DeleteTodo(ctx context.Context, req entity.DeleteTodoRequest) error
}

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type todoService struct {
	dbRW    repository.DatabaseReaderWriter
	cacheRW repository.CacheReaderWriter
}

func NewTodoService(
	dbRW repository.DatabaseReaderWriter,
	cacheRW repository.CacheReaderWriter,
) *todoService {
	return &todoService{
		dbRW:    dbRW,
		cacheRW: cacheRW,
	}
}
