package postgres

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_pgSQL_DeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta(`DELETE todos WHERE id = $1`)
	type args struct {
		ctx context.Context
		id  string
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
			args:  args{ctx: context.Background(), id: "01"},
			mocking: func(a args) {
				result := sqlmock.NewResult(0, 1)
				mock.ExpectExec(query).WithArgs(a.id).WillReturnResult(result)
			},
			wantErr: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			tc.mocking(tc.args)

			pg := pgSQL{db: db}
			err := pg.DeleteTodo(tc.args.ctx, tc.args.id)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
