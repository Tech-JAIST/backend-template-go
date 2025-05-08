package main

import (
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/db/migration"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// setup database
	db, err := gorm.Open(
		mysql.New(mysql.Config{DSNConfig: config.MySQL()}),
		config.GORM(),
	)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			e.Logger.Fatal(err)
		}
		sqlDB.Close()
	}()

	// migrate tables
	if err := migration.MigrateTables(db); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo)
	v1API := e.Group("/api/v1")
	h.SetupRoutes(v1API)

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
