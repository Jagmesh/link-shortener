package database

import (
	"fmt"
	"link-shortener/config"
	"link-shortener/pkg/logger"
	"reflect"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(config *config.DbConfig) *Database {
	db, err := gorm.Open(postgres.Open(makeDSN(*config)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	logger.GetLogger().Info("Connected to database")
	return &Database{db}
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
