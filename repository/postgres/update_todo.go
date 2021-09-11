package postgres

import (
	"context"
	"database/sql"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (p *pgSQL) UpdateTodo(ctx context.Context, todo entity.Todo) error {
	query := `UPDATE todos SET title = $2, description = $3, status = $4 WHERE id = $1`

	p.mu.Lock()
	res, err := p.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.Description, todo.Status)
	p.mu.Unlock()
	if err != nil {
		return err
	}

	if val, err := res.RowsAffected(); err != nil || val == 0 {
		return sql.ErrNoRows
	}

	return nil
}
