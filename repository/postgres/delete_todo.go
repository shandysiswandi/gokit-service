package postgres

import (
	"context"
	"database/sql"

	"github.com/shandysiswandi/gokit-service/pkg/logger"
)

func (p *pgSQL) DeleteTodo(ctx context.Context, id string) error {
	logger.Info("pgSQL.DeleteTodo")
	query := `DELETE todos WHERE id = $1`

	p.mu.Lock()
	res, err := p.db.ExecContext(ctx, query, id)
	p.mu.Unlock()
	if err != nil {
		logger.Error("pgSQL.GetAllTodo - db.ExecContext [err] " + err.Error())
		return err
	}

	if val, err := res.RowsAffected(); err != nil || val == 0 {
		logger.Error("pgSQL.GetAllTodo - data not found")
		return sql.ErrNoRows
	}

	return nil
}
