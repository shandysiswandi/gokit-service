package postgres

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/pkg/logger"
)

func (p *pgSQL) GetAllTodo(ctx context.Context) ([]entity.Todo, error) {
	logger.Info("pgSQL.GetAllTodo")
	query := `SELECT * FROM todos`

	p.mu.RLock()
	rows, err := p.db.QueryContext(ctx, query)
	p.mu.RUnlock()
	if err != nil {
		logger.Error("pgSQL.GetAllTodo - db.QueryContext [err] " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status); err != nil {
			logger.Error("pgSQL.GetAllTodo - rows.Scan [err] " + err.Error())
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		logger.Error("pgSQL.GetAllTodo - rows.Next [err] " + err.Error())
		return nil, err
	}

	return todos, nil
}
