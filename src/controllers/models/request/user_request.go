package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3,max=40"`
	Password string `json:"password" binding:"required,containsany=!@#$%&*,min=6,max=20"`
}
