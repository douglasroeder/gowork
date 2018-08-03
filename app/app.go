package app

import (
	"github.com/jinzhu/gorm"
)

const (
	appName    = "goWork"
	appVersion = "0.0.1"
)

// App contains all configuration, stateful service connections, etc
type App struct {
	Config *Config
	DB     *gorm.DB
}

// Close handles closing the app, terminating DB connections, etc
func (app *App) Close() {
	defer app.DB.Close()
}

// NewApp returns a new instance of App
func NewApp() *App {
	config, err := NewConfig()
	if err != nil {
		panic("Error loading configuration")
	}

	db := initDB(config)

	return &App{
		Config: config,
		DB:     db,
	}
}

func initDB(config *Config) *gorm.DB {
	dbName := "./" + config.AppName + "_" + config.Environment + ".db"
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("Error initialising DB")
	}

	return db
}
