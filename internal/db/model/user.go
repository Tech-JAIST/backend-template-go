package model

type User struct {
	ID   string `gorm:"type:char(36);not null;primaryKey"`
	Name string `gorm:"type:varchar(32)"`
}
