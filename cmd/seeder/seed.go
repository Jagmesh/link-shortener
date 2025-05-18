package main

import (
	"link-shortener/config"
	"link-shortener/pkg/database"
	"link-shortener/seed"
	"os"
	"path/filepath"
)

func main() {
	baseDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	sqlDir := filepath.Join(baseDir, "seed", "sql")

	conf := config.GetConfig()
	db := database.New(&conf.Db)

	seeder := seed.Seeder{Db: db, SqlScriptsDir: sqlDir}
	seeder.CreateBaseTestRecords()
}
