package postgres

import (
	"client/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

func MustConnect(cfg *config.Config) *bun.DB {
	connStr := "user=" + cfg.User + " dbname=" + cfg.DBname + " password=" + cfg.Password + " sslmode=disable"
	con, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(con, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}
