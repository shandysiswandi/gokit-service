package postgres

import (
	"context"
	"database/sql"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (p *pgSQL) CreateTodo(ctx context.Context, t entity.Todo) error {
	query := `INSERT INTO todos (title, description, status) VALUES ($1, $2, $3)`

	p.mu.Lock()
	res, err := p.db.ExecContext(ctx, query, t.Title, t.Description, t.Status)
	p.mu.Unlock()
	if err != nil {
		return err
	}

	if val, err := res.RowsAffected(); err != nil || val == 0 {
		return sql.ErrNoRows
	}

	return nil
}
