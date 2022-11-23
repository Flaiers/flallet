package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/flaiers/flallet/internal/entity"
)

func NewDatabaseConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateDatabase(dsn string) error {
	db := NewDatabaseConnection(dsn)
	db.Exec(`
		DO $$ BEGIN
    		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
        		CREATE TYPE transaction_type AS ENUM ('accrual', 'withdrawal', 'transfer');
    		END IF;
		END $$;
	`)
	db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_status') THEN
				CREATE TYPE transaction_status AS ENUM ('pending', 'approved', 'rejected');
			END IF;
		END $$;
	`)

	return db.AutoMigrate(&entity.User{}, &entity.Transaction{}, &entity.Reservation{})
}
