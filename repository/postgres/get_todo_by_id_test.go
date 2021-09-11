package postgres

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_pgSQL_GetTodoByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM todos WHERE id = $1`)
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
			title: "Negative_row.Scan",
			args:  args{ctx: context.Background()},
			mocking: func(a args) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("1")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			title: "Positive",
			args:  args{ctx: context.Background(), id: "001"},
			mocking: func(a args) {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "status"}).
					AddRow("1", "some title", "description one", "draft")
				mock.ExpectQuery(query).WithArgs(a.id).WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			tc.mocking(tc.args)

			pg := pgSQL{db: db}
			_, err := pg.GetTodoByID(tc.args.ctx, tc.args.id)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
