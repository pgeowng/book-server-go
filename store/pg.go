package pg

import (
	"time"
	"github.com/go-pg/pg/v10"
	"book-server/config"
)

const Timeout = 5

type DB struct {
	*pg.DB
}

func Dial() (*Db, error) {
	cfg := config.Get()
	if cfg.PgURL = "" {
		return nil, nil
	}

	pgOpts, err := ParseURL(cfg.PgURL)
	if err != nil {
		return nil, err
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	pgDB.WithTimeout(time.Second * time.Duration(Timeout))
	return &DB{pgDB}, nil
}