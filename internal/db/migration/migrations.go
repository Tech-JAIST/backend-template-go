package migration

import (
	"backend/internal/db/model"
	"errors"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func v1() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250508_add_root_user",
		Migrate: func(db *gorm.DB) error {
			if !db.Migrator().HasTable(&model.User{}) {
				return errors.New("user table not found")
			}

			return db.Create(&model.User{
				ID:   uuid.MustParse("0d9a2f49-99aa-49e6-bcf6-b123953aca63"),
				Name: "root",
			}).Error
		},
	}
}
