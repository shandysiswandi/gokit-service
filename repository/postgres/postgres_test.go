package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgres(t *testing.T) {
	testTable := []struct {
		title   string
		conf    func() Configuration
		wantErr bool
	}{
		{
			title: "Negative db.Ping",
			conf: func() Configuration {
				return Configuration{}
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		_, err := NewPostgres(tc.conf())
		assert.Equal(t, tc.wantErr, err != nil)
	}
}

func Test_pgSQL_GetDB(t *testing.T) {
	// This test does not use a test table because it has no parameters and only one or two lines of code.
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	pg := &pgSQL{db: db}
	dbi := pg.GetDB()
	assert.NotNil(t, dbi)
}

func Test_pgSQL_Close(t *testing.T) {
	// This test does not use a test table because it has no parameters and only one or two lines of code.
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectClose()
	defer mock.ExpectationsWereMet()

	pg := &pgSQL{db: db}
	err = pg.Close()
	assert.NoError(t, err)
}
