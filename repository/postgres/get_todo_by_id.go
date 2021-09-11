package postgres

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/pkg/logger"
)

func (p *pgSQL) GetTodoByID(ctx context.Context, id string) (entity.Todo, error) {
	logger.Info("pgSQL.GetTodoByID")
	query := `SELECT * FROM todos WHERE id = $1`

	p.mu.RLock()
	row := p.db.QueryRowContext(ctx, query, id)
	p.mu.RUnlock()

	var todo entity.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status); err != nil {
		logger.Error("pgSQL.GetTodoByID - rows.Scan [err] " + err.Error())
		return todo, err
	}

	return todo, nil
}
