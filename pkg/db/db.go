package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	LogLevel int
	MaxIdle  int
	MaxOpen  int
	MaxLife  time.Duration
}

func (p *PostgresOptions) GetDSN(pg *PostgresOptions) string {
	return "host=" + pg.Host + " port=" + string(pg.Port) + " user=" + pg.User + " password=" + pg.Password + " dbname=" + pg.DBName + " sslmode=disable"
}

func NewPostgresConnection(opts *PostgresOptions) (gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}
	db, err := gorm.Open(postgres.Open(opts.GetDSN(opts)), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(opts.MaxIdle)
	sqlDB.SetMaxOpenConns(opts.MaxOpen)
	sqlDB.SetConnMaxLifetime(opts.MaxLife)
	return *db, nil
}
