package postgres

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (p *pgSQL) GetAllTodo(ctx context.Context) ([]entity.Todo, error) {
	query := `SELECT * FROM todos`

	p.mu.RLock()
	rows, err := p.db.QueryContext(ctx, query)
	p.mu.RUnlock()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
