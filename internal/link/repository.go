package link

import (
	"errors"
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

func (r *Repository) Create(link *Link) (*Link, error) {
	err := r.Database.Create(link).Error
	if err != nil {
		return nil, err
	}
	return link, nil
}

type FindParams struct {
	hash string
	url  string
	id   uint
}

func (r *Repository) FindFirst(params *FindParams) (*Link, error) {
	var link Link
	if params.hash != "" {
		res := r.Database.First(&link, "hash = ?", params.hash)
		return &link, res.Error
	}
	if params.url != "" {
		res := r.Database.First(&link, "url = ?", params.url)
		return &link, res.Error
	}
	if params.id != 0 {
		res := r.Database.First(&link, "id = ?", params.id)
		return &link, res.Error
	}

	return nil, errors.New("no params provided")
}

func (r *Repository) Delete(link *Link) error {
	return r.Database.Delete(link).Error
}
