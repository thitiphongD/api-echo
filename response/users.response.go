package response

type ResponseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	Role     string `json:"role"`
}
