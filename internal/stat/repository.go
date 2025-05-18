package stat

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"link-shortener/entity/model"
	"link-shortener/pkg/database"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrIncrementClick(linkId uint) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}, {Name: "day_date"}},
		DoUpdates: clause.Assignments(map[string]any{
			"click": gorm.Expr("stats.click + 1"),
		}),
	}).Create(model.NewStat(linkId)).Error
}

func (r *Repository) GetClicksCountByDate(linkId []uint, from string, to string) (uint, error) {
	var totalClicksCount uint

	err := r.db.
		Model(&model.Stat{}).
		Select("SUM(click)").
		Table("stats").
		Where("link_id IN ?", linkId).
		Where("day_date > ?", from).
		Where("day_date < ?", to).
		Scan(&totalClicksCount).Error

	return totalClicksCount, err
}
