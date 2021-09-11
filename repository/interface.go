package repository

import (
	"context"
	"database/sql"

	"github.com/shandysiswandi/gokit-service/entity"
)

type DatabaseReaderWriter interface {
	GetAllTodo(ctx context.Context) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, id string) ([]entity.Todo, error)
	CreateTodo(ctx context.Context, todo entity.Todo) error
	UpdateTodo(ctx context.Context, todo entity.Todo) error
	DeleteTodo(ctx context.Context, id string) error
	//
	GetDB() *sql.DB
	Close() error
}

type CacheReaderWriter interface {
	//
	Close() error
}
