package models

type Film struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	FilmUrl     string `json:"film_url"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Hero        string `json:"hero"`
	Trailer     string `json:"trailer"`
}

func (Film) TableName() string {
	return "films"
}

// type Product struct {
// 	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
// 	Name       string               `json:"name" form:"name" gorm:"type: varchar(255)"`
// 	Desc       string               `json:"desc" gorm:"type:text" form:"desc"`
// 	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
// 	CategoryID []int                `json:"category_id" form:"category_id" gorm:"-"`
// 	Price      int                  `json:"price" form:"price" gorm:"type: int"`
// 	FilmUrl    string               `json:"film_url" form:"film"`
// 	Image      string               `json:"image" form:"image" gorm:"type: varchar(255)"`
// 	UserID     int                  `json:"user_id" form:"user_id"`
// 	User       UsersProfileResponse `json:"user"`
// 	CreatedAt  time.Time            `json:"-"`
// 	UpdatedAt  time.Time            `json:"-"`
// }

// type ProductResponse struct {
// 	ID         int                  `json:"id"`
// 	Name       string               `json:"name"`
// 	Desc       string               `json:"desc"`
// 	Price      int                  `json:"price"`
// 	Image      string               `json:"image"`
// 	FilmUrl    string               `json:"film_url" form:"film"`
// 	UserID     int                  `json:"-"`
// 	User       UsersProfileResponse `json:"user"`
// 	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
// 	CategoryID []int                `json:"category_id" form:"category_id" gorm:"-"`
// }

// type ProductUserResponse struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Desc   string `json:"desc"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	UserID int    `json:"-"`
// }

// func (ProductResponse) TableName() string {
// 	return "products"
// }

// func (ProductUserResponse) TableName() string {
// 	return "products"
// }
