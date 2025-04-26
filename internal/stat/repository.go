package stat

import (
	"link-shortener/entity/model"
	"link-shortener/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrIncrementClick(linkId uint) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"click": gorm.Expr("stats.click + 1"),
		}),
	}).Create(model.NewStat(linkId)).Error
}
