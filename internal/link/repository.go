package link

import (
	"link-shortener/model"
	"link-shortener/pkg/database"
)

type Repository struct {
	Database *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		Database: db,
	}
}

func (r *Repository) Create(link *model.Link) (*model.Link, error) {
	err := r.Database.Create(link).Error
	if err != nil {
		return nil, err
	}
	return link, nil
}

type FindParams struct {
	hash   string
	url    string
	id     uint
	userId uint
}

func (r *Repository) FindFirst(params *FindParams) (*model.Link, error) {
	var link model.Link
	query := r.Database.Model(&model.Link{})

	if params.hash != "" {
		query = query.Where("hash = ?", params.hash)
	}
	if params.url != "" {
		query = query.Where("url = ?", params.url)
	}
	if params.id != 0 {
		query = query.Where("id = ?", params.id)
	}
	if params.userId != 0 {
		query.Where("user_id = ?", params.userId)
	}

	if err := query.First(&link).Error; err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *Repository) Delete(link *model.Link) error {
	return r.Database.Delete(link, ).Error
}
