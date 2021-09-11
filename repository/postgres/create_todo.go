package postgres

import (
	"context"
	"database/sql"

	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/pkg/logger"
)

func (p *pgSQL) CreateTodo(ctx context.Context, t entity.Todo) error {
	logger.Info("pgSQL.CreateTodo")
	query := `INSERT INTO todos (title, description, status) VALUES ($1, $2, $3)`

	p.mu.Lock()
	res, err := p.db.ExecContext(ctx, query, t.Title, t.Description, t.Status)
	p.mu.Unlock()
	if err != nil {
		logger.Error("pgSQL.CreateTodo ExecContext " + err.Error())
		return err
	}

	if val, err := res.RowsAffected(); err != nil || val == 0 {
		logger.Error("pgSQL.CreateTodo - failed save to db")
		return sql.ErrNoRows
	}

	return nil
}
