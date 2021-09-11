package postgres

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_pgSQL_GetAllTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM todos`)
	type args struct {
		ctx context.Context
	}

	testTable := []struct {
		title   string
		args    args
		mocking func(a args)
		wantErr bool
	}{
		{
			title: "Negative - QueryContext",
			args:  args{ctx: context.Background()},
			mocking: func(a args) {
				mock.ExpectQuery(query).WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			title: "Negative_rows.Scan",
			args:  args{ctx: context.Background()},
			mocking: func(a args) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("1")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			title: "Positive",
			args:  args{ctx: context.Background()},
			mocking: func(a args) {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "status"}).
					AddRow("1", "some title", "description one", "draft")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			tc.mocking(tc.args)

			pg := pgSQL{db: db}
			_, err := pg.GetAllTodo(tc.args.ctx)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
