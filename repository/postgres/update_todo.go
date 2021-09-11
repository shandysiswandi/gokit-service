package postgres

import (
	"context"
	"database/sql"

	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/pkg/logger"
)

func (p *pgSQL) UpdateTodo(ctx context.Context, todo entity.Todo) error {
	logger.Info("pgSQL.UpdateTodo")
	query := `UPDATE todos SET title = $2, description = $3, status = $4 WHERE id = $1`

	p.mu.Lock()
	res, err := p.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.Description, todo.Status)
	p.mu.Unlock()
	if err != nil {
		logger.Error("pgSQL.UpdateTodo - db.ExecContext [err] " + err.Error())
		return err
	}

	if val, err := res.RowsAffected(); err != nil || val == 0 {
		logger.Error("pgSQL.UpdateTodo - data not found")
		return sql.ErrNoRows
	}

	return nil
}
