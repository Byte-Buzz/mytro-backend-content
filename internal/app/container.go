package app

import (
	"mytro-backend-content/internal/infrastructure/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB

	Services     *AppServices
	repositories *AppRepositories
}

type AppServices struct {
}

type AppRepositories struct {
}

func NewApp(config *config.Config, db *gorm.DB, logger *zap.Logger) *App {
	app := &App{
		Config: config,
		Logger: logger,
		DB:     db,

		Services:     &AppServices{},
		repositories: &AppRepositories{},
	}

	app.createRepositories()
	app.createServices()

	return app
}

func (app *App) createRepositories() {
}

func (app *App) createServices() {
}
