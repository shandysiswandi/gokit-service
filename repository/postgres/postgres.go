package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

// pgSQL
type pgSQL struct {
	db *sql.DB
	tx *sql.Tx
	mu sync.RWMutex
}

type Configuration struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Options  string
}

// NewPostgres
func NewPostgres(conf Configuration) (*pgSQL, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Options,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &pgSQL{db: db, tx: nil}, nil
}

func (p *pgSQL) Close() error {
	return p.db.Close()
}

func (p *pgSQL) GetDB() *sql.DB {
	return p.db
}
