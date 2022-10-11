package db

import (
	"database/sql"
	"fmt"
	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
	"web-server/pkg/config"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	if db == nil {
		configs := config.GetConfig()

		d, err := sql.Open(configs.DBDriver, configs.DBSource)
		if err != nil {
			return nil, fmt.Errorf("sql.Open failed: %w", err)
		}
		db = d
	}

	return db, nil
}

func MigrateDb() error {
	db, err := InitDB()
	if err != nil {
		return fmt.Errorf("InitDB failed: %v", err)
	}

	migrator, err := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "../../pkg/db/migrations")
	if err != nil {
		return fmt.Errorf("NewMigrator failed: %v", err)
	}
	return migrator.Migrate()
}
