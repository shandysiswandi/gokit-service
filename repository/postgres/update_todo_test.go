package postgres

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/stretchr/testify/assert"
)

func Test_pgSQL_UpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta(`UPDATE todos SET title = $2, description = $3, status = $4 WHERE id = $1`)
	type args struct {
		ctx  context.Context
		todo entity.Todo
	}

	testTable := []struct {
		title   string
		args    args
		mocking func(a args)
		wantErr bool
	}{
		{
			title: "Negative ExecContext",
			args:  args{ctx: context.Background()},
			mocking: func(a args) {
				mock.ExpectExec(query).WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			title: "Negative RowsAffected",
			args:  args{ctx: context.Background()},
			mocking: func(a args) {
				result := sqlmock.NewResult(0, 0)
				mock.ExpectExec(query).WillReturnResult(result)
			},
			wantErr: true,
		},
		{
			title: "Positive",
			args:  args{ctx: context.Background(), todo: entity.Todo{ID: "11", Title: "aa", Description: "bb", Status: "cc"}},
			mocking: func(a args) {
				result := sqlmock.NewResult(0, 1)
				mock.ExpectExec(query).
					WithArgs(a.todo.ID, a.todo.Title, a.todo.Description, a.todo.Status).
					WillReturnResult(result)
			},
			wantErr: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			tc.mocking(tc.args)

			pg := pgSQL{db: db}
			err := pg.UpdateTodo(tc.args.ctx, tc.args.todo)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
