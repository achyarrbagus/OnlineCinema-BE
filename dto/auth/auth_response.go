package authdto

type LoginResponse struct {
	Email  string `json:"email" form:"email" validate:"required"`
	Token  string `json:"token" validate:"required"`
	Role   string `json:"role"`
	UserID int    `json:"userId"`
}
