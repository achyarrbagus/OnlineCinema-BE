package repositories

import (
	"backEnd/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	CreateFilm(film models.Film) (models.Film, error)
	FindFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	GetOneFilm(ID, User int) (models.Film, error)
	DeleteFilm(film models.Film) (models.Film, error)
	UpdateFilm(Film models.Film) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
// 	var products models.Product
// 	err := r.db.Delete(&product).Error // Using Delete method

// 	return products, err
// }

func (r *repository) UpdateFilm(Film models.Film) (models.Film, error) {
	err := r.db.Save(&Film).Error

	return Film, err
}

func (r *repository) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error

	return film, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error
	return film, err
}

func (r *repository) FindFilm() ([]models.Film, error) {
	var Film []models.Film
	err := r.db.Raw("SELECT * FROM films").Scan(&Film).Error

	return Film, err
}

func (r *repository) GetOneFilm(ID, User int) (models.Film, error) {
	var film models.Film
	err := r.db.First(&film, ID).Error

	var trans models.Transaction
	error := r.db.Where("film_id = ? AND user_id = ?", ID, User).First(&trans).Error
	if error == nil && trans.Status == "success" {
		film.Price = 0
	}

	return film, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.First(&film, ID).Error

	return film, err
}
