package repositories

import (
	"backEnd/models"
	"net/url"

	// "strconv"

	"gorm.io/gorm"
)

type FilterRepository interface {
	MultiFilter(params url.Values) ([]models.Film, error)
}

func RepositoryFilter(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) MultiFilter(params url.Values) ([]models.Film, error) {
	var films []models.Film

	category := params.Get("category")
	err := r.db.Where("category = ?", category).Find(&films).Error

	return films, err

}
