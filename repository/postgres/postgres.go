package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/shandysiswandi/gokit-service/pkg/logger"
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
	logger.Info("NewPostgres")
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
		logger.Error("NewPostgres sql.Open " + err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error("NewPostgres db.Ping " + err.Error())
		return nil, err
	}

	logger.Error("NewPostgres successfully connected")
	return &pgSQL{db: db, tx: nil}, nil
}

func (p *pgSQL) Close() error {
	logger.Error("NewPostgres close connection")
	return p.db.Close()
}

func (p *pgSQL) GetDB() *sql.DB {
	logger.Error("NewPostgres get db connection instance")
	return p.db
}
