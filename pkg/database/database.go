package database

import (
	"fmt"
	"link-shortener/config"
	"link-shortener/pkg/logger"
	"reflect"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var retries uint8 = 0

func New(config *config.DbConfig) *Database {
	db, err := gorm.Open(postgres.Open(makeDSN(config.Credentials)), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return reconnect(err, config)
	}
	logger.GetLogger().Info("Connected to database")
	return &Database{db}
}

func reconnect(err error, config *config.DbConfig) *Database {
	if retries >= config.MaxRetriesNumber {
		panic(fmt.Sprintf("DB connection failed: %s. Max retries (%d) reached", err, retries))
	}
	retries++

	logger.GetLogger().Infof("Failed to connect to database. Attemp #%d of %d. Retrying ...", retries, config.MaxRetriesNumber)
	time.Sleep(getExponentialDelay(retries))
	return New(config)
}

func (db *Database) Migrate(models ...any) {
	logger.GetLogger().Debug("Running migrations...")
	if err := db.AutoMigrate(models...); err != nil {
		panic(err)
	}
	logger.GetLogger().Debug("Migrations completed")
}

func makeDSN(params any) string {
	var result []string

	v := reflect.ValueOf(params)

	t := v.Type()
	for i := range v.NumField() {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		result = append(result, fmt.Sprintf("%s=%v", strings.ToLower(key), value))
	}

	return strings.Join(result, " ")
}
