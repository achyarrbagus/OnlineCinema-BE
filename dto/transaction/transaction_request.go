package transactiondto

import (
	"backEnd/models"
)

type CreateTransactionRequest struct {
	Status    string                      `json:"status"`
	Title     string                      `json:"title"`
	OrderDate string                      `json:"order_date"`
	Price     int                         `json:"price"`
	UserID    int                         `json:"user_id" form:"user_id"`
	User      models.UsersProfileResponse `json:"user"`
	FilmID    int                         `json:"film_id" form:"film_id"`
	Film      models.Film                 `json:"film"`
}
