package migration

import (
	"backend/internal/db/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) error {
	err:=db.AutoMigrate(InitialTables()...)
	if err != nil {
		return err
	}
	return gormigrate.New(db, gormigrate.DefaultOptions, Migrations()).Migrate()
}

// InitialTables returns the initial tables to be created in the database.
func InitialTables() []interface{} {
	return []interface{}{
		&model.User{},
	}
}

// MigrateTables returns the migrations to be run on the database.
// If you want to add columns to existing tables, you can do so by adding a new migration.
func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		v1(),
	}
}
