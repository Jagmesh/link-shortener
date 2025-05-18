package seed

import (
	"link-shortener/pkg/database"
	"link-shortener/pkg/logger"
	"os"
	"path/filepath"
)

var log = logger.GetWithScopes("SEEDER")

type Seeder struct {
	Db            *database.Database
	SqlScriptsDir string
}

func (s *Seeder) CreateBaseTestRecords() {
	dirEntries, err := os.ReadDir(s.SqlScriptsDir)
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		log.Infof("Processing file: '%s'...", entry.Name())

		file, err := os.ReadFile(filepath.Join(s.SqlScriptsDir, entry.Name()))
		if err != nil {
			panic(err)
		}

		if err := s.Db.Exec(string(file)).Error; err != nil {
			panic(err)
		}
		log.Infof("✅ Succesfully processed file '%s!\n", entry.Name())
	}

}
