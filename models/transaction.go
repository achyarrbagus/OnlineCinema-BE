package models

type Transaction struct {
	ID        int                  `json:"id"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"foreignKey:UserID"`
	FilmID    int                  `json:"film_id"`
	Film      Film                 `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Status    string               `json:"status"`
	Title     string               `json:"title"`
	OrderDate string               `json:"order_date"`
	Price     int                  `json:"price"`
}

func (Transaction) TableName() string {
	return "transactions"
}
