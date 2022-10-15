package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/config"
)

func NewDBClient(ctx context.Context, env *config.Env) (*sqlx.DB, error) {
	c := mysql.Config{
		DBName:               env.MySqlDB,
		User:                 env.MySqlUser,
		Passwd:               env.MySqlPass,
		Addr:                 fmt.Sprintf("%s:%s", env.MySqlHost, env.MySqlPort),
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
		Loc:                  time.FixedZone("Asia/Tokyo", 9*60*60),
	}

	fmt.Println(env.MySqlHost)
	fmt.Println(c.Addr)

	dsn := c.FormatDSN()
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db client (%s): %w", dsn, err)
	}

	db.SetConnMaxLifetime(time.Second * time.Duration(env.MySQLConnMaxLifetime))

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database %s: %w", dsn, err)
	}

	// https://github.com/jmoiron/sqlx/issues/143
	db = db.Unsafe()
	return db, nil
}
