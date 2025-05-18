package link

import (
	"link-shortener/entity/model"
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
	Hash   string
	Url    string
	Id     uint
	UserId uint
}

func (r *Repository) FindAll(params *FindParams) ([]model.Link, error) {
	var links []model.Link
	query := r.Database.Model(&model.Link{})

	if params.Hash != "" {
		query = query.Where("hash = ?", params.Hash)
	}
	if params.Url != "" {
		query = query.Where("url = ?", params.Url)
	}
	if params.Id != 0 {
		query = query.Where("id = ?", params.Id)
	}
	if params.UserId != 0 {
		query.Where("user_id = ?", params.UserId)
	}

	err := query.Find(&links).Error
	return links, err
}

func (r *Repository) FindOne(params *FindParams) (*model.Link, error) {
	var link model.Link
	query := r.Database.Model(&model.Link{})

	if params.Hash != "" {
		query = query.Where("hash = ?", params.Hash)
	}
	if params.Url != "" {
		query = query.Where("url = ?", params.Url)
	}
	if params.Id != 0 {
		query = query.Where("id = ?", params.Id)
	}
	if params.UserId != 0 {
		query.Where("user_id = ?", params.UserId)
	}

	err := query.First(&link).Error
	return &link, err
}

func (r *Repository) Delete(link *model.Link) error {
	return r.Database.Delete(link).Error
}
