package filmdto

type FilmRequest struct {
	Title       string `json:"title" form:"name" validate:"required"`
	Category    string `json:"category" form:"category" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	FilmUrl     string `json:"film_url" form:"film_url" validate:"required"`
	Description string `json:"description" form:"desc" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"photo" validate:"required"`
	Trailer     string `json:"trailer" form:"trailer" validate:"required"`
	Hero        string `json:"hero" form:"hero" validate:"required"`
}
