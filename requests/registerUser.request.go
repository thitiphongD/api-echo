package requests

type RequestRegisterUser struct {
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
