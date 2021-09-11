package postgres

import (
	"context"
	"database/sql"
)

func (p *pgSQL) DeleteTodo(ctx context.Context, id string) error {
	query := `DELETE todos WHERE id = $1`

	p.mu.Lock()
	res, err := p.db.ExecContext(ctx, query, id)
	p.mu.Unlock()
	if err != nil {
		return err
	}

	if val, err := res.RowsAffected(); err != nil || val == 0 {
		return sql.ErrNoRows
	}

	return nil
}
